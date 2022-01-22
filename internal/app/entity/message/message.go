package ChatMessage

import "fmt"

type Message struct {
	ID         int64  `json:"id"`
	Chat       int64  `json:"chat"`
	Author     int64  `json:"author"`
	Text       string `json:"text"`
	Created_at string `json:"create_at"`
}

func (message Message) ValidateMessageData() error {
	if message.Chat <= 0 {
		return fmt.Errorf("chatID is not valid. Actually: %v", message.Chat)
	}

	if message.Author <= 0 {
		return fmt.Errorf("author is not valid. Actually: %v", message.Author)
	}

	if len(message.Text) == 0 {
		return fmt.Errorf("message text cannot be empty")
	}

	return nil
}
