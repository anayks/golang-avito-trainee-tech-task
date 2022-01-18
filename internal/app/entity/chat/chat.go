package ChatEntity

import (
	"time"

	ChatMessage "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/message"
)

type Chat struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Users      []int64   `json:"users"`
	Created_at time.Time `json:"created_at"`
}

func GetUserChatList(userID int64) (ChatsList []Chat) {
	return ChatsList
}

func GetMessageList(chatID int64) (MessagesList []ChatMessage.Message) {
	return MessagesList
}
