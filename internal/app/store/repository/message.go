package sqlstore

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	ChatEntity "github.com/anayks/golang-avito-tech-test/internal/app/entity/chat"
	ChatMessage "github.com/anayks/golang-avito-tech-test/internal/app/entity/message"
)

type RepositoryMessages struct {
	store *Store
}

func (r RepositoryMessages) Create(ctx context.Context, message *ChatMessage.Message) (id int64, err error) {

	tx, err := r.store.db.BeginTx(ctx, nil)
	defer tx.Rollback()

	result := int64(0)

	err = tx.QueryRowContext(ctx, "SELECT id from chatsUsers WHERE user_id = $1 and chat_id = $2", message.Author, message.Chat).Scan(&result)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("user or chat not found")
		}
		return 0, err
	}

	err = tx.QueryRowContext(ctx, "INSERT into messages (user_id, chat_id, text) VALUES ($1, $2, $3) RETURNING id", message.Author, message.Chat, message.Text).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("internal server error while creating message: %v", err)
	}

	if err = tx.Commit(); err != nil {
		return 0, fmt.Errorf("internal server error while creating message: %v", err)
	}

	return id, nil
}

func (r RepositoryMessages) GetChatMessages(chat *ChatEntity.Chat) (result string, err error) {
	fmt.Printf("Chat ID: %v", chat.ID)

	queryResult, err := r.store.db.Query("SELECT * FROM messages WHERE chat_id = $1", chat.ID)

	if err != nil {
		return "", err
	}

	defer queryResult.Close()

	var arrayResult []ChatMessage.Message

	for queryResult.Next() {
		fmt.Printf("qu ID: %v", chat.ID)

		queryItem := &ChatMessage.Message{}

		if err := queryResult.Scan(&queryItem.ID, &queryItem.Author, &queryItem.Chat, &queryItem.Text, &queryItem.Created_at); err != nil {
			bytesResult, err := json.Marshal(arrayResult)
			if err != nil {
				return result, err
			}

			result = string(bytesResult)

			return result, nil
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

	result = string(bytesResult)

	return result, nil
}
