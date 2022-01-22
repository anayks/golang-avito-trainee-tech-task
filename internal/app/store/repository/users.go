package sqlstore

import (
	"fmt"

	ChatUser "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/user"
)

type RepositoryUsers struct {
	store *Store
}

func (r RepositoryUsers) Create(user *ChatUser.ChatUser) (int64, error) {
	fmt.Printf("username: %v\n", user.Username)

	var id int64

	err := r.store.db.QueryRow("INSERT into users (username) VALUES ($1) RETURNING id", user.Username).Scan(&id)

	if err != nil {
		fmt.Printf("error creating user: %v", err)
		return 0, err
	}

	return id, nil
}
