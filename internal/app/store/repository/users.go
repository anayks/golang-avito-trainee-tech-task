package sqlstore

import (
	"fmt"

	chatUser "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/user"
)

type RepositoryUsers struct {
	store *Store
}

const (
	queryCreateUser = "INSERT into users (username) VALUES ($1) RETURNING id, created_at"
)

func (r RepositoryUsers) Create(username string) (chatUser.ChatUser, error) {
	chatUser := &chatUser.ChatUser{
		Username: username,
	}

	err := r.store.db.QueryRow(queryCreateUser, username).Scan(&chatUser.ID, &chatUser.Created_at)

	if err != nil {
		return *chatUser, fmt.Errorf("error creating user: %w", err)
	}

	return *chatUser, nil
}
