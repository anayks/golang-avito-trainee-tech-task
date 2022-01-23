package sqlstore

import (
	"fmt"

	chatUser "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/user"
)

type RepositoryUsers struct {
	store *Store
}

const (
	queryCreateUser = "INSERT into users (username) VALUES ($1) RETURNING id"
)

func (r RepositoryUsers) Create(user *chatUser.ChatUser) (int64, error) {

	var id int64

	err := r.store.db.QueryRow(queryCreateUser, user.Username).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error creating user: %w", err)
	}

	return id, nil
}
