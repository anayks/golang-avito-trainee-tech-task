package chatUser

import (
	"fmt"
	"regexp"
	"time"
)

type ChatUser struct {
	ID         int64  `json:"user"`
	Username   string `json:"username"`
	Created_at time.Time
}

func (user ChatUser) ValidateUserName() error {
	matched, err := regexp.Match(`^[a-zA-Z0-9а-яА-Я_ё]{4,20}$`, []byte(user.Username))

	if err != nil {
		return err
	}

	if !matched {
		return fmt.Errorf("invalid user name")
	}

	return nil
}

func (user ChatUser) ValidateUserID() error {
	if user.ID <= 0 {
		return fmt.Errorf("user is not valid. Actually: %v", user.ID)
	}
	return nil
}
