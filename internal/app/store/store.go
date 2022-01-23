package db

import (
	"database/sql"
	"fmt"
	"os"

	chatEntity "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/chat"
	chatMessage "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/message"
	chatUser "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/user"
	_ "github.com/lib/pq"
)

type RepositoryUsers interface {
	Create(*chatUser.ChatUser) int
}

type RepositoryChats interface {
	Create(*chatEntity.Chat) (int, error)
	GetUserChats(*chatUser.ChatUser) ([]chatEntity.Chat, error)
}

type RepositoryMessages interface {
	Create(*chatMessage.Message) (int, error)
	GetChatMessages(*chatEntity.Chat) (string, error)
}

type Store interface {
	Users() *RepositoryUsers
	Chats() *RepositoryChats
	Messages() *RepositoryMessages
}

func Connect() (database *sql.DB) {
	host := os.Getenv("DB_HOST")     // "db"
	port := os.Getenv("DB_PORT")     // 5432
	user := os.Getenv("DB_USER")     // "postgres"
	password := os.Getenv("DB_PASS") // "HEYO"
	dbname := os.Getenv("DB_NAME")   //"postgres"

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)

	CheckError(err)

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	return db
}

func CheckError(err error) {
	if err != nil {
		fmt.Printf("mysql is not connected. Error reason: %v", err)
		panic(err)
	}
}
