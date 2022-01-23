package sqlstore

import (
	"context"
	"database/sql"
	"encoding/json"
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

func (r RepositoryMessages) Create(ctx context.Context, message *chatMessage.Message) (id int64, err error) {

	tx, err := r.store.db.BeginTx(ctx, nil)
	defer tx.Rollback()

	var result int64

	err = tx.QueryRowContext(ctx, querySelectUserLinks, message.Author, message.Chat).Scan(&result)

	if err != nil && err != sql.ErrNoRows {
		return 0, fmt.Errorf("user or chat not found")
	}

	if err != nil {
		return 0, err
	}

	err = tx.QueryRowContext(ctx, queryCreateMessage, message.Author, message.Chat, message.Text).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("internal server error while creating message: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return 0, fmt.Errorf("internal server error while creating message: %w", err)
	}

	return id, nil
}

func (r RepositoryMessages) GetChatMessages(chat *chatEntity.Chat) (string, error) {
	var result string

	queryResult, err := r.store.db.Query(queryGetChatMessages, chat.ID)

	if err != nil {
		return "", err
	}

	defer queryResult.Close()

	var arrayResult []chatMessage.Message

	for queryResult.Next() {
		queryItem := &chatMessage.Message{}

		if err := queryResult.Scan(&queryItem.ID, &queryItem.Author, &queryItem.Chat, &queryItem.Text, &queryItem.Created_at); err != nil {
			bytesResult, err := json.Marshal(arrayResult)
			if err != nil {
				return result, err
			}

			return string(bytesResult), nil
		}

		arrayResult = append(arrayResult, *queryItem)
	}

	if err := queryResult.Err(); err != nil {
		return "", err
	}

	bytesResult, err := json.Marshal(arrayResult)

	if err != nil {
		return result, err
	}

	return string(bytesResult), nil
}
