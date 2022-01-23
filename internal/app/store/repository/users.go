package sqlstore

import (
	"fmt"

	ChatUser "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/user"
)

type RepositoryUsers struct {
	store *Store
}

func (r RepositoryUsers) Create(user *ChatUser.ChatUser) (int64, error) {

	var id int64

	err := r.store.db.QueryRow("INSERT into users (username) VALUES ($1) RETURNING id", user.Username).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error creating user: %w", err)
	}

	return id, nil
}
