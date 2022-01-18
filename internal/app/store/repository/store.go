package sqlstore

import (
	"database/sql"
)

type Store struct {
	db                 *sql.DB
	repositoryUsers    *RepositoryUsers
	repositoryChats    *RepositoryChats
	repositoryMessages *RepositoryMessages
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Users() RepositoryUsers {
	if s.repositoryUsers != nil {
		return *s.repositoryUsers
	}

	s.repositoryUsers = &RepositoryUsers{
		store: s,
	}

	return *s.repositoryUsers
}

func (s *Store) Chats() RepositoryChats {
	if s.repositoryChats != nil {
		return *s.repositoryChats
	}

	s.repositoryChats = &RepositoryChats{
		store: s,
	}

	return *s.repositoryChats
}

func (s *Store) Messages() RepositoryMessages {
	if s.repositoryMessages != nil {
		return *s.repositoryMessages
	}

	s.repositoryMessages = &RepositoryMessages{
		store: s,
	}

	return *s.repositoryMessages
}
