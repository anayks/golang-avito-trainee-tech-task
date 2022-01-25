package sqlstore

import (
	"context"
	"database/sql"
	"fmt"

	chatEntity "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/chat"
	chatMessage "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/message"
)

type RepositoryMessages struct {
	store *Store
}

const (
	querySelectUserLinks = "SELECT id from chatsUsers WHERE user_id = $1 and chat_id = $2"
	queryCreateMessage   = "INSERT into messages (user_id, chat_id, text) VALUES ($1, $2, $3) RETURNING id"
	queryGetChatMessages = "SELECT * FROM messages WHERE chat_id = $1"
)

func (r RepositoryMessages) Create(ctx context.Context, message *chatMessage.Message) (newMessage *chatMessage.Message, err error) {
	tx, err := r.store.db.BeginTx(ctx, nil)
	defer tx.Rollback()

	newMessage = &chatMessage.Message{
		Author: message.Author,
		Chat:   message.Chat,
		Text:   message.Text,
	}

	if err != nil {
		return newMessage, err
	}

	var result int64

	err = tx.QueryRowContext(ctx, querySelectUserLinks, message.Author, message.Chat).Scan(&result)

	if len(message.Text) > 320 {
		return newMessage, fmt.Errorf("message so long")
	}

	if err != nil && err != sql.ErrNoRows {
		return newMessage, fmt.Errorf("user or chat not found")
	}

	if err != nil {
		return newMessage, fmt.Errorf("error while select users: %w", err)
	}

	err = tx.QueryRowContext(ctx, queryCreateMessage, message.Author, message.Chat, message.Text).Scan(&newMessage.ID)

	if err != nil {
		return newMessage, fmt.Errorf("internal server error while creating message: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return newMessage, fmt.Errorf("internal server error while creating message: %w", err)
	}

	return newMessage, nil
}

func (r RepositoryMessages) GetChatMessages(chat *chatEntity.Chat) ([]chatMessage.Message, error) {
	queryResult, err := r.store.db.Query(queryGetChatMessages, chat.ID)

	if err != nil {
		return nil, err
	}

	defer queryResult.Close()

	var arrayResult []chatMessage.Message

	for queryResult.Next() {
		queryItem := &chatMessage.Message{}

		if err := queryResult.Scan(&queryItem.ID, &queryItem.Author, &queryItem.Chat, &queryItem.Text, &queryItem.Created_at); err != nil {
			return arrayResult, err
		}

		arrayResult = append(arrayResult, *queryItem)
	}

	if err := queryResult.Err(); err != nil {
		return nil, err
	}

	return arrayResult, nil
}
