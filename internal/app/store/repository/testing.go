package sqlstore

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	chatUser "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/user"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // ...
	"github.com/stretchr/testify/assert"
)

// TestDB ...
func TestDB(t *testing.T, relativePath string) (*sql.DB, func(...string)) {
	t.Helper()

	err := godotenv.Load(relativePath)

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		t.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			for _, table := range tables {
				db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", table))
			}
		}
		db.Close()
	}
}

func DBFiller(t *testing.T, db *sql.DB, teardown func(...string)) (chat_id int64, newUser *chatUser.ChatUser, message_id int64, teardownFiller func()) {
	newUser = &chatUser.ChatUser{}
	newUser.Username = "hey1"
	err := db.QueryRow("INSERT INTO users (username) VALUES ($1) RETURNING ID, created_at", newUser.Username).Scan(&newUser.ID, &newUser.Created_at)
	assert.NoError(t, err)

	err = db.QueryRow("INSERT INTO chats (chatname) VALUES ('hey1') RETURNING ID").Scan(&chat_id)
	assert.NoError(t, err)

	_, err = db.Exec("INSERT INTO chatsUsers (chat_id, user_id) VALUES ($1, $2) RETURNING ID", chat_id, newUser.ID)
	assert.NoError(t, err)

	teardownFiller = func() {
		teardown("users")
		teardown("chatsUsers")
		teardown("chats")
		teardown("messages")
	}

	err = db.QueryRow("INSERT INTO messages (chat_id, user_id, text) VALUES ($1, $2, $3) RETURNING ID", chat_id, newUser.ID, "heyo!").Scan(&message_id)
	assert.NoError(t, err)

	return chat_id, newUser, message_id, teardownFiller
}
