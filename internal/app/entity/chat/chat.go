package ChatEntity

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

func (chat *Chat) VaildateChatData() error {
	if chat.ID < 0 {
		return fmt.Errorf("chat ID is not valid")
	}

	matched, err := regexp.Match(`^[a-zA-Z0-9а-яА-Я_]{4,20}$`, []byte(chat.Name))
	if err != nil {
		return err
	}

	if !matched {
		return fmt.Errorf("invalid chat name characters. Actually: %s", chat.Name)
	}

	return nil
}
