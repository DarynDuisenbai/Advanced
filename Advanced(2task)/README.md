# Chat App with MongoDB

## Description

The project is a web application for real-time messaging. It is developed using the Golang programming language, Gorilla WebSocket library, and MongoDB as the database.

## Group

- SE -2214

## Contributors

- Daryn
- Asylzhan

## Screenshot

![Chat App](C:\Users\Admin\GolandProjects\chatapp\screenshot.png)

## Running the Application

1. Install Golang and MongoDB.
2. Clone the repository: `git clone https://github.com/yourusername/chat-app-mongodb.git`
3. Navigate to the project folder: `cd chat-app-mongodb`
4. Apply migrations: `migrate -path db/migrations -database "mongodb://username:password@localhost:27017/chatdb" up`
5. Run the server: `go run main.go`
6. Open the web page in your browser: `http://localhost:8081`

## Tools Used

- Golang
- Gorilla WebSocket
- MongoDB
- golang-migrate

## Source Links

- [Gorilla WebSocket](https://github.com/gorilla/websocket)
- [MongoDB Go Driver](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo)
- [golang-migrate](https://github.com/golang-migrate/migrate)

