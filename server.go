package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"leaderboard/pkg/leaderboardDB"
	"log"
	"net"
	"net/mail"
	"os"
	"time"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/metadata"

	pb "leaderboard/proto"

	"google.golang.org/grpc"
)

const port = ":8080"

type LeaderboardServer struct {
	pb.UnimplementedLeaderboardServer
}

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       uint64 `json:"id"`
	jwt.StandardClaims
}

// createToken creates and returns a new token.
func createToken(jwtKey string, username string, email string, id uint64) string {
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		Username: username,
		Email:    email,
		ID:       id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		log.Fatal(err)
	}

	return tokenString
}

// checkToken checks if the provided token is valid.
func checkToken(jwtKey string, ctx context.Context) error {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("no metadata found in context")
	}

	tokens := headers.Get("token")

	if len(tokens) < 1 {
		return fmt.Errorf("no token found in metadata")
	}

	var userClaim Claims

	token, err := jwt.ParseWithClaims(tokens[0], &userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return fmt.Errorf("invalid token")
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

// Registration registers new users and returns a token if registration is successfully completed.
func (s *LeaderboardServer) Registration(ctx context.Context, in *pb.RegistrationRequest) (*pb.RegistrationResponse, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE")))
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	username := in.GetUsername()
	email := in.GetEmail()
	password := in.GetPassword()
	confirmPassword := in.GetConfirmPassword()

	var id uint64

	if len(username) < 1 {
		return nil, fmt.Errorf("too short username: enter at least 1 symbol")
	}

	_, err = mail.ParseAddress(email)

	if err != nil {
		return nil, fmt.Errorf("invalid email")
	}
	if len(password) < 6 {
		return nil, fmt.Errorf("too short password: enter at least 6 symbols")
	}
	if password != confirmPassword {
		return nil, fmt.Errorf("password is not confirmed")
	}

	var exists int

	rows, err := db.Query(fmt.Sprintf("SELECT EXISTS(SELECT * FROM `users` WHERE `username`='%s');", username))
	if err != nil {
		log.Fatal(err)
	}
	rows.Next()
	if err := rows.Scan(&exists); err != nil {
		log.Fatal(err)
	}
	if exists == 1 {
		return nil, fmt.Errorf("user with this username already exists")
	}
	rows.Close()

	rows, err = db.Query(fmt.Sprintf("SELECT EXISTS(SELECT * FROM `users` WHERE `email`='%s');", email))
	if err != nil {
		log.Fatal(err)
	}
	rows.Next()
	if err := rows.Scan(&exists); err != nil {
		log.Fatal(err)
	}
	if exists == 1 {
		return nil, fmt.Errorf("user with this email already exists")
	}
	rows.Close()

	_, err = db.Query(fmt.Sprintf("INSERT INTO `users` (`username`, `email`, `password`) VALUES ('%s', '%s', '%s');", username, email, password))
	if err != nil {
		log.Fatal(err)
	}

	rows, err = db.Query(fmt.Sprintf("SELECT `ID` FROM `users` WHERE `username`='%s';", username))
	if err != nil {
		log.Fatal(err)
	}
	rows.Next()
	if err := rows.Scan(&id); err != nil {
		log.Fatal(err)
	}
	rows.Close()

	jwtKey := os.Getenv("AUTH_KEY")
	tokenString := createToken(jwtKey, username, email, id)

	return &pb.RegistrationResponse{Token: tokenString}, nil
}

// LogIn checks if the user is logged in and returns a token if the check is successfully completed.
func (s *LeaderboardServer) LogIn(ctx context.Context, in *pb.LogInRequest) (*pb.LogInResponse, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE")))
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	var (
		username string
		email    string
		password string
		id       uint64
	)

	username = in.GetUsername()
	password = in.GetPassword()

	rows, err := db.Query(fmt.Sprintf("SELECT `email`, `id` FROM `users` WHERE `username`='%s' AND `password`='%s';", username, password))
	if err != nil {
		log.Fatal(err)
	}
	if !rows.Next() {
		return nil, fmt.Errorf("incorrect username or password")
	}
	if err := rows.Scan(&email, &id); err != nil {
		log.Fatal(err)
	}
	rows.Close()

	jwtKey := os.Getenv("AUTH_KEY")
	tokenString := createToken(jwtKey, username, email, id)

	return &pb.LogInResponse{Token: tokenString}, nil
}

// Find returns the rank of player with given name. Find works with a stream connection, so all the requests and responses are performed within one
// persistent stream.
func (s *LeaderboardServer) Find(stream pb.Leaderboard_FindServer) error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// Token
	ctx := stream.Context()
	jwtKey := os.Getenv("AUTH_KEY")

	err = checkToken(jwtKey, ctx)
	if err != nil {
		return err
	}
	//

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE")))
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	var (
		name  string
		score uint32
	)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		name = req.GetName()
		score = req.GetScore()

		rank, err := leaderboardDB.GetRank(db, name, score)

		if err := stream.Send(&pb.RankResponse{Rank: rank}); err != nil {
			return err
		}
	}
	return nil
}

// GetLeaderboard fetches the leaderboard.
func (s *LeaderboardServer) GetLeaderboard(ctx context.Context, in *pb.LeaderboardRequest) (*pb.LeaderboardResponse, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// Token
	jwtKey := os.Getenv("AUTH_KEY")

	err = checkToken(jwtKey, ctx)
	if err != nil {
		return &pb.LeaderboardResponse{}, err
	}
	//

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE")))
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	name := in.GetName()
	page := in.GetPage()
	monthly := in.GetMonthly()
	resAmount := in.GetResAmount()
	r, a, next_page, err := leaderboardDB.Fetch(db, name, page, monthly, resAmount)
	if err != nil {
		return &pb.LeaderboardResponse{}, err
	}
	results := make([]*pb.Player, len(r))
	around_me := make([]*pb.Player, len(a))
	for i := 0; i < len(r); i++ {
		results[i] = &pb.Player{Name: r[i].Name, Score: r[i].Score, Rank: r[i].Rank}

	}
	for i := 0; i < len(a); i++ {
		around_me[i] = &pb.Player{Name: a[i].Name, Score: a[i].Score, Rank: a[i].Rank}
	}
	return &pb.LeaderboardResponse{Results: results, AroundMe: around_me, NextPage: next_page}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLeaderboardServer(s, &LeaderboardServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve: %v", err)
	}
}
