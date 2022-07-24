# About the leaderboard project

 This is a leaderboard API for games. The communication with the API made via gRPC. The API server written in GoLang.

# Deployment and testing

 Download the project:
 ```bash
 git clone https://github.com/Matfej28/leaderboard
  ```
 Open downloaded directory, open the Docker and compose the file:
 ```bash
 docker compose up
  ```
Wait until everything is loaded, then open the API testing platform (like Postman) and test the API. To provide your token, add Metadata with the key "token" and insert your token there as a value.
