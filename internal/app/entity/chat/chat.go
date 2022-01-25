package chatEntity

import (
	"fmt"
	"regexp"
	"time"
)

type Chat struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Users      []int64   `json:"users"`
	Created_at time.Time `json:"created_at"`
}

func (chat *Chat) ValidateChatData() error {

	matched, err := regexp.Match(`^[a-zA-Z0-9а-яА-Я_ё]{4,20}$`, []byte(chat.Name))
	if err != nil {
		return err
	}

	if !matched {
		return fmt.Errorf("invalid chat name characters. Actually: %s", chat.Name)
	}

	return nil
}
