package db

import (
	"database/sql"
	"fmt"

	ChatEntity "github.com/anayks/golang-avito-tech-test/internal/app/entity/chat"
	ChatMessage "github.com/anayks/golang-avito-tech-test/internal/app/entity/message"
	ChatUser "github.com/anayks/golang-avito-tech-test/internal/app/entity/user"
	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "HEYO"
	dbname   = "postgres"
)

type RepositoryUsers interface {
	Create(*ChatUser.ChatUser) int
}

type RepositoryChats interface {
	Create(*ChatEntity.Chat) (int, error)
	GetUserChats(*ChatUser.ChatUser) ([]ChatEntity.Chat, error)
}

type RepositoryMessages interface {
	Create(*ChatMessage.Message) (int, error)
	GetChatMessages(*ChatEntity.Chat) (string, error)
}

type Store interface {
	Users() *RepositoryUsers
	Chats() *RepositoryChats
	Messages() *RepositoryMessages
}

func Connect() (database *sql.DB) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

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
