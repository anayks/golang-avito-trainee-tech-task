package sqlstore_test

import (
	"context"
	"testing"

	chatEntity "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/chat"
	chatMessage "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/message"
	sqlstore "github.com/anayks/golang-avito-trainee-tech-task/internal/app/store/repository"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryMessages_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, "../../../../.env")
	defer db.Close()

	s := sqlstore.New(db)

	chatID, user, _, teardownFiller := sqlstore.DBFiller(t, db, teardown)
	defer teardownFiller()

	t.Run("TestRepositoryMessages_Create", func(t *testing.T) {
		message := chatMessage.TestObject()
		message.Author = user.ID
		message.Chat = chatID
		message.Text = "Hello!"

		id, err := s.RepositoryMessages.Create(context.Background(), message)
		assert.NoError(t, err)
		assert.NotEqual(t, 0, id)
	})
}

func TestRepositoryMessages_GetChatMessages(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, "../../../../.env")
	defer db.Close()

	s := sqlstore.New(db)

	chatID, _, messageID, teardownFiller := sqlstore.DBFiller(t, db, teardown)
	defer teardownFiller()

	t.Run("TestRepositoryMessages_GetChatMessages", func(t *testing.T) {
		chat := chatEntity.TestObject()
		chat.ID = chatID
		result, err := s.RepositoryMessages.GetChatMessages(chat)
		assert.NoError(t, err)
		assert.Equal(t, messageID, result[0].ID)
	})
}
