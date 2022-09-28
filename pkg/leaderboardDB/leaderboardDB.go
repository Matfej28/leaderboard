package leaderboardDB

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Player struct {
	Name  string
	Score uint32
	Rank  uint32
}

// calculateRank calculates the rank of the player with given name and score.
func calculateRank(db *sql.DB, name string, score uint32) uint32 {
	var newRank = uint32(1)

	rows, err := db.Query("SELECT `name`, `score`, `rank`  FROM `leaderboard` ORDER BY `Rank` DESC;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		n string
		s uint32
		r uint32
	)

	for rows.Next() {
		if err := rows.Scan(&n, &s, &r); err != nil {
			log.Fatal(err)
		}
		if score > s {
			_, err = db.Query(fmt.Sprintf("UPDATE `leaderboard` SET `Rank`= %d WHERE `Name` = '%s';", r+1, n))
			if err != nil {
				log.Fatal(err)
			}
			newRank = r
		} else {
			newRank = r + 1
			break
		}
	}
	year, month, _ := time.Now().Date()
	_, err = db.Query(fmt.Sprintf("INSERT INTO `leaderboard` VALUES ('%s', %d, %d, %d, %d);", name, score, newRank, year, month))
	if err != nil {
		log.Fatal(err)
	}
	return newRank
}

// GetRank inserts the player with given name and score into the db and returns players' rank in the leaderboard. If the given score is not larger than the
// players' best score, GetRank just returns the rank of the player. If the given score is the new best score, GetRank updates the score in the db and returns
// the new rank of the player.
func GetRank(db *sql.DB, name string, score uint32) (rank uint32, globErr error) {
	if name == "" {
		globErr = errors.New("invalid name, write at least 1 character")
		return
	}

	rows, err := db.Query(fmt.Sprintf("SELECT `name`, `score`, `rank` FROM `leaderboard` WHERE `Name`='%s';", name))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		n string
		s uint32
		r uint32
	)

	for rows.Next() {
		if err := rows.Scan(&n, &s, &r); err != nil {
			log.Fatal(err)
		}
		if score > s {
			_, err := db.Query(fmt.Sprintf("DELETE FROM `leaderboard` WHERE `Name` = '%s';", name))
			if err != nil {
				log.Fatal(err)
			}
			recalculateLeaderboard(db)
		} else {
			return r, nil
		}
	}
	return calculateRank(db, name, score), nil
}

// Fetch returns a page of results, which consists of resAmount results and which number is "page". It is possible to see monthly and all time leaderboard.
// Every returned result consists of players' name, score and position(rank). If the name of the player is passed and the player is not in this list of results
// (and their result is not in any of the previous pages), a list of players around the current player will be returned.
func Fetch(db *sql.DB, name string, page uint32, monthly bool, resAmount uint32) (results []Player, around_me []Player, next_page uint32, globErr error) {
	if resAmount == 0 {
		resAmount = 10
	}
	var appears bool

	if page == 0 {
		page = 1
	}

	next_page = page + 1

	var skip = uint32(resAmount) * (page - 1)
	var counter = resAmount
	var (
		n string
		s uint32
		r uint32
		y int
		m int
	)
	var tmp Player

	if monthly {
		var rankCounter uint32 = resAmount * (page - 1)

		year, month, _ := time.Now().Date()
		rows, err := db.Query(fmt.Sprintf("SELECT * FROM `leaderboard` WHERE `Year`=%d AND `Month`=%d ORDER BY `Score` DESC;", year, month))
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for counter > 0 && rows.Next() {
			if skip > 0 {
				skip--
				continue
			}
			if err := rows.Scan(&n, &s, &r, &y, &m); err != nil {
				log.Fatal(err)
			}
			rankCounter++
			results = append(results, Player{n, s, rankCounter})
			if n == name {
				appears = true
			}
			counter--
		}

		tmp = Player{n, s, rankCounter}

		if !rows.Next() {
			if len(results) == 0 {
				globErr = errors.New("page number is invalid")
			}
			next_page = 0
			return
		}
		rankCounter++
		if err := rows.Scan(&n, &s, &r, &y, &m); err != nil {
			log.Fatal(err)
		}
		if appears {
			return
		}

		counter = 3 //amount of players around to return

		for counter > 0 {
			if appears {
				around_me = append(around_me, Player{n, s, rankCounter})
				counter--
			}
			if n == name {
				around_me = append(around_me, tmp)
				around_me = append(around_me, Player{n, s, rankCounter})
				appears = true
			}
			tmp = Player{n, s, rankCounter}
			if !rows.Next() {
				break
			}
			rankCounter++
			if err := rows.Scan(&n, &s, &r, &y, &m); err != nil {
				log.Fatal(err)
			}
		}
	} else {
		rows, err := db.Query("SELECT * FROM `leaderboard` ORDER BY `Rank`;")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for counter > 0 && rows.Next() {
			if skip > 0 {
				skip--
				continue
			}
			if err := rows.Scan(&n, &s, &r, &y, &m); err != nil {
				log.Fatal(err)
			}
			results = append(results, Player{n, s, r})
			if n == name {
				appears = true
			}
			counter--
		}

		tmp = Player{n, s, r}

		if !rows.Next() {
			if len(results) == 0 {
				globErr = errors.New("page number is invalid")
			}
			next_page = 0
			return
		}
		if err := rows.Scan(&n, &s, &r, &y, &m); err != nil {
			log.Fatal(err)
		}
		if appears {
			return
		}

		counter = 3 //amount of players around to return

		for counter > 0 {
			if appears {
				around_me = append(around_me, Player{n, s, r})
				counter--
			}
			if n == name {
				around_me = append(around_me, tmp)
				around_me = append(around_me, Player{n, s, r})
				appears = true
			}
			tmp = Player{n, s, r}
			if !rows.Next() {
				break
			}
			if err := rows.Scan(&n, &s, &r, &y, &m); err != nil {
				log.Fatal(err)
			}
		}
	}
	return
}

// recalculateLeaderboard recalculates the rank of every player in db.
func recalculateLeaderboard(db *sql.DB) {
	var counter uint32

	rows, err := db.Query("SELECT `name`, `score`, `rank` FROM `leaderboard` ORDER BY `Rank`;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var (
		n string
		s uint32
		r uint32
	)
	for rows.Next() {
		counter++
		if err := rows.Scan(&n, &s, &r); err != nil {
			log.Fatal(err)
		}
		if counter == r {
			continue
		}
		_, err = db.Query(fmt.Sprintf("UPDATE `leaderboard` SET `Rank`= %d WHERE `Name` = '%s';", r-1, n))
		if err != nil {
			log.Fatal(err)
		}
	}
}
