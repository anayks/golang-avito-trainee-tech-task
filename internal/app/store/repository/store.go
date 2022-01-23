package sqlstore

import (
	"database/sql"
)

type Store struct {
	db                 *sql.DB
	RepositoryUsers    *RepositoryUsers
	RepositoryChats    *RepositoryChats
	RepositoryMessages *RepositoryMessages
}

func New(db *sql.DB) *Store {
	newStore := &Store{
		db: db,
	}

	newStore.RepositoryChats = &RepositoryChats{
		store: newStore,
	}

	newStore.RepositoryMessages = &RepositoryMessages{
		store: newStore,
	}

	newStore.RepositoryUsers = &RepositoryUsers{
		store: newStore,
	}

	return newStore
}
