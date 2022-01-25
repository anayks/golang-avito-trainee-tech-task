package chatMessage_test

import (
	"fmt"
	"testing"

	chatMessage "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/message"
	"github.com/stretchr/testify/assert"
)

func TestMessage_ValidateMessageData(t *testing.T) {
	testCases := []struct {
		toMessage func() *chatMessage.Message
		isValid   bool
	}{
		{
			toMessage: func() *chatMessage.Message {
				obj := chatMessage.TestObject()
				obj.Author = 25
				obj.Chat = 25
				obj.Text = ""
				return obj
			},
			isValid: false,
		},
		{
			toMessage: func() *chatMessage.Message {
				obj := chatMessage.TestObject()
				obj.Author = 25
				obj.Chat = 25
				obj.Text = "3"
				return obj
			},
			isValid: true,
		},
		{
			toMessage: func() *chatMessage.Message {
				obj := chatMessage.TestObject()
				obj.Author = 0
				obj.Chat = 25
				obj.Text = "3"
				return obj
			},
			isValid: false,
		},
		{
			toMessage: func() *chatMessage.Message {
				obj := chatMessage.TestObject()
				obj.Author = 355
				obj.Chat = 0
				obj.Text = "5"
				return obj
			},
			isValid: false,
		},
		{
			toMessage: func() *chatMessage.Message {
				obj := chatMessage.TestObject()
				obj.Author = 0
				obj.ID = 0
				obj.Text = ""
				return obj
			},
			isValid: false,
		},
		{
			toMessage: func() *chatMessage.Message {
				obj := chatMessage.TestObject()
				obj.Author = 35
				obj.ID = 1
				obj.Text = "Heyo!"
				return obj
			},
			isValid: true,
		},
	}

	for index, val := range testCases {
		testCaseName := fmt.Sprintf("%s %d", "message", index)
		t.Run(testCaseName, func(t *testing.T) {
			if val.isValid == (val.toMessage().ValidateMessageData() == nil) {
				assert.NoError(t, nil)
			} else {
				assert.Error(t, val.toMessage().ValidateMessageData())
			}
		})
	}
}
