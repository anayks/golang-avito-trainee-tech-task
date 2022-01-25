package sqlstore

import (
	"context"
	"database/sql"
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

func (r RepositoryChats) Create(ctx context.Context, chat *chatEntity.Chat) (newChat *chatEntity.Chat, err error) {
	tx, err := r.store.db.BeginTx(ctx, nil)
	defer tx.Rollback()

	newChat = &chatEntity.Chat{
		Users: chat.Users,
		Name:  chat.Name,
	}

	if err != nil {
		return newChat, err
	}

	err = tx.QueryRowContext(ctx, "INSERT into chats (chatname) VALUES ($1) RETURNING id", chat.Name).Scan(&newChat.ID)

	if err != nil && err == sql.ErrNoRows {
		return newChat, fmt.Errorf("inserting chat went wrong: %w", err)
	}

	if err != nil {
		return newChat, err
	}

	for _, val := range chat.Users {
		err := CreateChatUserLinksTx(ctx, tx, r, newChat.ID, val)
		if err == nil {
			continue
		}

		return newChat, fmt.Errorf("(chat_id %v) not exists and error: %w", newChat.ID, err)
	}

	if err = tx.Commit(); err != nil {
		return newChat, err
	}

	return newChat, nil
}

func (r RepositoryChats) GetUserChats(user *chatUser.ChatUser) ([]chatEntity.Chat, error) {
	rows, err := r.store.db.Query(querySelectAllChatInfo, user.ID)

	if err != nil && err == sql.ErrNoRows {
		return nil, fmt.Errorf("user from id %v is not in some chat", user.ID)
	}

	defer rows.Close()

	var resultArray []chatEntity.Chat

	for rows.Next() {
		chatResult := &chatEntity.Chat{}

		if err := rows.Scan(&chatResult.ID, pq.Array(&chatResult.Users), &chatResult.Name, &chatResult.Created_at); err != nil {
			return nil, err
		}
		resultArray = append(resultArray, *chatResult)
	}

	return resultArray, nil
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
