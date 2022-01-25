package chatEntity_test

import (
	"fmt"
	"testing"

	chatEntity "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/chat"
	"github.com/stretchr/testify/assert"
)

func TestChat_ValidateChatData(t *testing.T) {
	testCases := []struct {
		toChat  func() *chatEntity.Chat
		isValid bool
	}{
		{
			toChat: func() *chatEntity.Chat {
				chat := chatEntity.TestObject()
				chat.Name = "приветяё"
				chat.ID = 25
				return chat
			},
			isValid: true,
		},
		{
			toChat: func() *chatEntity.Chat {
				chat := chatEntity.TestObject()
				chat.Name = "№;%:№;:"
				chat.ID = 1
				return chat
			},
			isValid: false,
		},
		{
			toChat: func() *chatEntity.Chat {
				chat := chatEntity.TestObject()
				chat.Name = "привет"
				chat.ID = 0
				return chat
			},
			isValid: true,
		},
		{
			toChat: func() *chatEntity.Chat {
				chat := chatEntity.TestObject()
				chat.Name = "hellobuddy"
				chat.ID = 1
				return chat
			},
			isValid: true,
		},
		{
			toChat: func() *chatEntity.Chat {
				chat := chatEntity.TestObject()
				chat.Name = "2%@%@%@"
				chat.ID = 0
				return chat
			},
			isValid: false,
		},
		{
			toChat: func() *chatEntity.Chat {
				chat := chatEntity.TestObject()
				chat.Name = "hellobuddy_привет"
				chat.ID = 1
				return chat
			},
			isValid: true,
		},
		{
			toChat: func() *chatEntity.Chat {
				chat := chatEntity.TestObject()
				chat.Name = "hellobuddy_hellobuddy_hellobuddy_hellobuddy_hellobuddy_"
				chat.ID = 25
				return chat
			},
			isValid: false,
		},
		{
			toChat: func() *chatEntity.Chat {
				chat := chatEntity.TestObject()
				chat.Name = "hellobuddy_hellobuddy_hellobuddy_hellobuddy_hellobuddy_"
				chat.ID = -1
				return chat
			},
			isValid: false,
		},
		{
			toChat: func() *chatEntity.Chat {
				chat := chatEntity.TestObject()
				chat.Name = "3"
				chat.ID = 25
				return chat
			},
			isValid: false,
		},
	}

	for index, val := range testCases {
		testCaseName := fmt.Sprintf("%s %d", "chat", index)
		t.Run(testCaseName, func(t *testing.T) {
			if val.isValid {
				assert.NoError(t, val.toChat().ValidateChatData())
			} else {
				assert.Error(t, val.toChat().ValidateChatData())
			}
		})
	}
}
