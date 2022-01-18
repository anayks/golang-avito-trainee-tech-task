package sqlstore

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	ChatEntity "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/chat"
	ChatUser "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/user"
	pq "github.com/lib/pq"
)

type RepositoryChats struct {
	store *Store
}

func checkUserExistsWithTx(ctx context.Context, tx *sql.Tx, r RepositoryChats, chat_id int64, user_id int64) error {
	_, err := tx.ExecContext(ctx, "INSERT into chatsUsers (chat_id, user_id) VALUES ($1, $2)", chat_id, user_id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}
	return nil
}

func (r RepositoryChats) Create(ctx context.Context, chat *ChatEntity.Chat) (id int64, err error) {
	tx, err := r.store.db.BeginTx(ctx, nil)

	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	if err := tx.QueryRowContext(ctx, "INSERT into chats (chatname) VALUES ($1) RETURNING id", chat.Name).Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("inserting chat went wrong")
		}
		return 0, err
	}

	for _, val := range chat.Users {
		err := checkUserExistsWithTx(ctx, tx, r, id, val)
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

func (r RepositoryChats) GetUserChats(user *ChatUser.ChatUser) (chatsList string, err error) {
	rows, err := r.store.db.Query(
		`WITH res AS(
			SELECT 
				chatsUsers.chat_id AS res_id, 
				chats.chatname AS chatname, 
				chats.created_at AS created_at 
			FROM 
				chats 
				LEFT JOIN chatsUsers ON chats.id = chatsUsers.chat_id 
			WHERE 
				chatsUsers.user_id = $1 
			ORDER BY 
				user_id asc
		), 
		chatRelatives AS(
			SELECT 
				chat_id, 
				array_agg(chatsUsers.user_id) users 
			from 
				chatsUsers 
			WHERE 
				chat_id IN (
					SELECT 
						res_id 
					FROM 
						res
				) 
			GROUP BY 
				chatsUsers.chat_id
		) 
		SELECT 
			chat_id, 
			users, 
			chatname, 
			created_at 
		from 
			chatRelatives 
			INNER JOIN res ON chatRelatives.chat_id = res.res_id
		`, user.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			return chatsList, fmt.Errorf("user from id %v is not in some chat", user.ID)
		}
		return chatsList, err
	}

	defer rows.Close()

	var resultArray []ChatEntity.Chat

	for rows.Next() {
		chatResult := &ChatEntity.Chat{}

		if err := rows.Scan(&chatResult.ID, pq.Array(&chatResult.Users), &chatResult.Name, &chatResult.Created_at); err != nil {
			return chatsList, err
		}
		resultArray = append(resultArray, *chatResult)
	}

	byteResult, err := json.Marshal(resultArray)

	if err != nil {
		return "", err
	}

	chatsList = string(byteResult)

	return chatsList, nil
}
