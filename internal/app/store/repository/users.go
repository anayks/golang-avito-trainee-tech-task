package sqlstore

import (
	"fmt"

	ChatUser "github.com/anayks/golang-avito-tech-test/internal/app/entity/user"
)

type RepositoryUsers struct {
	store *Store
}

func (r RepositoryUsers) Create(user *ChatUser.ChatUser) (id int64) {
	fmt.Printf("username: %v\n", user.Username)

	err := r.store.db.QueryRow("INSERT into users (username) VALUES ($1) RETURNING id", user.Username).Scan(&id)

	if err != nil {
		fmt.Printf("error creating user: %v", err)
		return 0
	}

	return id
}
