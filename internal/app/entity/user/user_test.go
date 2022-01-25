package chatUser_test

import (
	"fmt"
	"testing"

	chatUser "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/user"
	"github.com/stretchr/testify/assert"
)

func TestUser_ValidateUserName(t *testing.T) {
	testCases := []struct {
		toUser  func() *chatUser.ChatUser
		isValid bool
	}{
		{
			toUser: func() *chatUser.ChatUser {
				chat := chatUser.TestObject()
				chat.Username = "приветяё"
				chat.ID = 25
				return chat
			},
			isValid: true,
		},
		{
			toUser: func() *chatUser.ChatUser {
				chat := chatUser.TestObject()
				chat.Username = "№;%:№;:"
				chat.ID = 1
				return chat
			},
			isValid: false,
		},
		{
			toUser: func() *chatUser.ChatUser {
				chat := chatUser.TestObject()
				chat.Username = "привет"
				chat.ID = 0
				return chat
			},
			isValid: true,
		},
		{
			toUser: func() *chatUser.ChatUser {
				chat := chatUser.TestObject()
				chat.Username = "hellobuddy"
				chat.ID = 1
				return chat
			},
			isValid: true,
		},
		{
			toUser: func() *chatUser.ChatUser {
				chat := chatUser.TestObject()
				chat.Username = "2%@%@%@"
				chat.ID = 0
				return chat
			},
			isValid: false,
		},
		{
			toUser: func() *chatUser.ChatUser {
				chat := chatUser.TestObject()
				chat.Username = "hellobuddy_привет"
				chat.ID = 1
				return chat
			},
			isValid: true,
		},
		{
			toUser: func() *chatUser.ChatUser {
				chat := chatUser.TestObject()
				chat.Username = "hellobuddy_hellobuddy_hellobuddy_hellobuddy_hellobuddy_"
				chat.ID = 25
				return chat
			},
			isValid: false,
		},
		{
			toUser: func() *chatUser.ChatUser {
				chat := chatUser.TestObject()
				chat.Username = "hellobuddy_hellobuddy_hellobuddy_hellobuddy_hellobuddy_"
				chat.ID = -1
				return chat
			},
			isValid: false,
		},
		{
			toUser: func() *chatUser.ChatUser {
				chat := chatUser.TestObject()
				chat.Username = "3"
				chat.ID = 25
				return chat
			},
			isValid: false,
		},
	}

	for index, val := range testCases {
		testCaseName := fmt.Sprintf("%s %d", "user", index)
		t.Run(testCaseName, func(t *testing.T) {
			if val.isValid == (val.toUser().ValidateUserName() == nil) {
				assert.NoError(t, nil)
			} else {
				assert.Error(t, val.toUser().ValidateUserName())
			}
		})
	}
}

func TestUser_ValidateUserID(t *testing.T) {
	testCases := []struct {
		toUser  func() *chatUser.ChatUser
		isValid bool
	}{
		{
			toUser: func() *chatUser.ChatUser {
				chat := chatUser.TestObject()
				chat.Username = "приветяё"
				chat.ID = 25
				return chat
			},
			isValid: true,
		},
		{
			toUser: func() *chatUser.ChatUser {
				chat := chatUser.TestObject()
				chat.Username = "№;%:№;:"
				chat.ID = 1
				return chat
			},
			isValid: true,
		},
		{
			toUser: func() *chatUser.ChatUser {
				chat := chatUser.TestObject()
				chat.Username = "привет"
				chat.ID = 0
				return chat
			},
			isValid: false,
		},
		{
			toUser: func() *chatUser.ChatUser {
				chat := chatUser.TestObject()
				chat.Username = "hellobuddy"
				chat.ID = 1
				return chat
			},
			isValid: true,
		},
		{
			toUser: func() *chatUser.ChatUser {
				chat := chatUser.TestObject()
				chat.Username = "2%@%@%@"
				chat.ID = 0
				return chat
			},
			isValid: false,
		},
		{
			toUser: func() *chatUser.ChatUser {
				chat := chatUser.TestObject()
				chat.Username = "hellobuddy_привет"
				chat.ID = 1
				return chat
			},
			isValid: true,
		},
		{
			toUser: func() *chatUser.ChatUser {
				chat := chatUser.TestObject()
				chat.Username = "hellobuddy_hellobuddy_hellobuddy_hellobuddy_hellobuddy_"
				chat.ID = 25
				return chat
			},
			isValid: true,
		},
		{
			toUser: func() *chatUser.ChatUser {
				chat := chatUser.TestObject()
				chat.Username = "hellobuddy_hellobuddy_hellobuddy_hellobuddy_hellobuddy_"
				chat.ID = -1
				return chat
			},
			isValid: false,
		},
		{
			toUser: func() *chatUser.ChatUser {
				chat := chatUser.TestObject()
				chat.Username = "3"
				chat.ID = 25
				return chat
			},
			isValid: true,
		},
	}

	for index, val := range testCases {
		testCaseName := fmt.Sprintf("%s %d", "user", index)
		t.Run(testCaseName, func(t *testing.T) {
			if val.isValid == (val.toUser().ValidateUserID() == nil) {
				assert.NoError(t, nil)
			} else {
				assert.Error(t, val.toUser().ValidateUserID())
			}
		})
	}
}
