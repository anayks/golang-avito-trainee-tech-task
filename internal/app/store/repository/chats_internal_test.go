package sqlstore_test

import (
	"context"
	"testing"

	chatEntity "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/chat"
	chatUser "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/user"
	sqlstore "github.com/anayks/golang-avito-trainee-tech-task/internal/app/store/repository"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryChats_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, "../../../../.env")
	defer db.Close()

	_, user, _, teardownFiller := sqlstore.DBFiller(t, db, teardown)
	defer teardownFiller()

	ctx := context.Background()

	s := sqlstore.New(db)
	e := chatEntity.TestObject()

	e.Name = "heyo"
	e.Users = append(e.Users, user.ID)

	result, err := s.RepositoryChats.Create(ctx, e)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestRepositoryChats_GetUserChats(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, "../../../../.env")
	defer teardown("chatsUsers")
	defer teardown("chats")

	chatID, user, _, teardownFiller := sqlstore.DBFiller(t, db, teardown)
	defer teardownFiller()

	t.Run("TestRepositoryChats_GetUserChats", func(t *testing.T) {
		s := sqlstore.New(db)
		e := chatUser.TestObject()
		e.ID = user.ID
		result, err := s.RepositoryChats.GetUserChats(e)
		assert.NoError(t, err)
		assert.Equal(t, chatID, result[0].ID)
	})
}
