package sqlstore

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	chatEntity "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/chat"
	chatUser "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/user"
	pq "github.com/lib/pq"
)

const (
	querySelectAllChatInfo = `WITH t_user_chats AS (
		SELECT 
			chat_id 
		FROM 
			chatsUsers
		WHERE
			user_id = $1
	)
	SELECT 
			chats.id AS id,
			ARRAY_AGG(DISTINCT chatsUsers.user_id) users,
			chats.chatname,
			chats.created_at
	FROM 
			chatsUsers
	JOIN t_user_chats ON t_user_chats.chat_id = chatsUsers.chat_id
	JOIN chats ON chats.id = t_user_chats.chat_id
	INNER JOIN
			messages messages1
	ON
			messages1.chat_id = chatsUsers.chat_id 
	AND
			t_user_chats.chat_id = messages1.chat_id
	AND
			messages1.created_at = (
				SELECT 
					DISTINCT MAX(messages.created_at)
				FROM 
					messages
				WHERE
					messages.chat_id = t_user_chats.chat_id
				)
	GROUP BY
			chats.id,
			chats.chatname,
			messages1.created_at,
			t_user_chats.chat_id
	ORDER BY 
			messages1.created_at DESC`
	queryCreateUserChatLinks = `INSERT into chatsUsers (chat_id, user_id) VALUES ($1, $2)`
)

type RepositoryChats struct {
	store *Store
}

func (r RepositoryChats) Create(ctx context.Context, chat *chatEntity.Chat) (int64, error) {
	tx, err := r.store.db.BeginTx(ctx, nil)

	var id int64

	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	err = tx.QueryRowContext(ctx, "INSERT into chats (chatname) VALUES ($1) RETURNING id", chat.Name).Scan(&id)

	if err != nil && err == sql.ErrNoRows {
		return 0, fmt.Errorf("inserting chat went wrong: %w", err)
	}

	if err != nil {
		return 0, err
	}

	for _, val := range chat.Users {
		err := CreateChatUserLinksTx(ctx, tx, r, id, val)
		if err == nil {
			continue
		}

		return 0, fmt.Errorf("(user_id %v) not exists", id)
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return id, nil
}

func (r RepositoryChats) GetUserChats(user *chatUser.ChatUser) (string, error) {
	rows, err := r.store.db.Query(querySelectAllChatInfo, user.ID)

	var chatsList string

	if err != nil && err == sql.ErrNoRows {
		return chatsList, fmt.Errorf("user from id %v is not in some chat", user.ID)
	}

	defer rows.Close()

	var resultArray []chatEntity.Chat

	for rows.Next() {
		chatResult := &chatEntity.Chat{}

		if err := rows.Scan(&chatResult.ID, pq.Array(&chatResult.Users), &chatResult.Name, &chatResult.Created_at); err != nil {
			return chatsList, err
		}
		resultArray = append(resultArray, *chatResult)
	}

	byteResult, err := json.Marshal(resultArray)

	if err != nil {
		return "", err
	}

	return string(byteResult), nil
}

func CreateChatUserLinksTx(ctx context.Context, tx *sql.Tx, r RepositoryChats, chat_id int64, user_id int64) error {
	_, err := tx.ExecContext(ctx, queryCreateUserChatLinks, chat_id, user_id)

	if err != nil && err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}
