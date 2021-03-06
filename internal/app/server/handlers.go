package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	chatEntity "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/chat"
	chatMessage "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/message"
	user "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/user"
)

func (s *server) handleAddUser(rw http.ResponseWriter, r *http.Request) {
	parsedUser := &user.ChatUser{}

	if err := json.NewDecoder(r.Body).Decode(parsedUser); err != nil {
		s.error(rw, r, http.StatusBadRequest, err)
		return
	}

	if err := parsedUser.ValidateUserName(); err != nil {
		s.error(rw, r, http.StatusBadRequest, err)
		return
	}

	user, err := s.store.RepositoryUsers.Create(parsedUser.Username)

	if err != nil {
		s.error(rw, r, http.StatusInternalServerError, err)
		return
	}

	response, err := json.Marshal(map[string]interface{}{
		"id": user.ID,
	})

	if err != nil {
		s.error(rw, r, http.StatusInternalServerError, err)
		return
	}

	s.respond(rw, r, http.StatusOK, string(response))
}

func (s *server) handlerCreateChat(rw http.ResponseWriter, r *http.Request) {
	parsedChat := &chatEntity.Chat{}

	if err := json.NewDecoder(r.Body).Decode(&parsedChat); err != nil {
		s.error(rw, r, http.StatusBadRequest, err)
		return
	}

	if err := parsedChat.ValidateChatData(); err != nil {
		s.error(rw, r, http.StatusBadRequest, fmt.Errorf("internal error while creating chat"))
		return
	}

	chat, err := s.store.RepositoryChats.Create(r.Context(), parsedChat)

	if err != nil {
		s.error(rw, r, http.StatusInternalServerError, fmt.Errorf("internal error while creating chat"))
		return
	}

	response, err := json.Marshal(map[string]interface{}{
		"id": chat.ID,
	})

	if err != nil {
		s.error(rw, r, http.StatusInternalServerError, err)
		return
	}

	s.respond(rw, r, http.StatusOK, string(response))
}

func (s *server) handlerSendMessage(rw http.ResponseWriter, r *http.Request) {
	parsedMessage := &chatMessage.Message{}

	if err := json.NewDecoder(r.Body).Decode(&parsedMessage); err != nil {
		s.error(rw, r, http.StatusBadRequest, err)
		return
	}

	if err := parsedMessage.ValidateMessageData(); err != nil {
		s.error(rw, r, http.StatusBadRequest, fmt.Errorf("internal server error while creating message"))
		return
	}

	message, err := s.store.RepositoryMessages.Create(r.Context(), parsedMessage)

	if err != nil {
		s.error(rw, r, http.StatusInternalServerError, fmt.Errorf("internal server error while creating message"))
		return
	}

	response, err := json.Marshal(map[string]interface{}{
		"id": message.ID,
	})

	if err != nil {
		s.error(rw, r, http.StatusInternalServerError, err)
		return
	}

	s.respond(rw, r, http.StatusOK, string(response))
}

func (s *server) handlerGetUserListOfChats() http.HandlerFunc {
	parsedUser := &user.ChatUser{}

	return func(rw http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&parsedUser); err != nil {
			s.error(rw, r, http.StatusBadRequest, err)
			return
		}

		if err := parsedUser.ValidateUserID(); err != nil {
			s.error(rw, r, http.StatusBadRequest, err)
			return
		}

		result, err := s.store.RepositoryChats.GetUserChats(parsedUser)

		if err != nil {
			s.error(rw, r, http.StatusInternalServerError, fmt.Errorf("internal server error while getting user's chats"))
			return
		}

		jsonResult, err := json.Marshal(result)

		if err != nil {
			s.error(rw, r, http.StatusInternalServerError, fmt.Errorf("internal server error while parsing user's chats"))
			return
		}

		s.respond(rw, r, http.StatusOK, string(jsonResult))
	}
}

func (s *server) handlerGetChatMessages() http.HandlerFunc {
	type parseServer struct {
		ID int64 `json:"chat"`
	}

	return func(rw http.ResponseWriter, r *http.Request) {
		parsedChat := &parseServer{}

		if err := json.NewDecoder(r.Body).Decode(&parsedChat); err != nil {
			s.error(rw, r, http.StatusBadRequest, err)
			return
		}

		if parsedChat.ID <= 0 {
			s.error(rw, r, http.StatusBadRequest, fmt.Errorf("chat-id is not valid. Actually: %v", parsedChat.ID))
			return
		}

		chatEntity := &chatEntity.Chat{
			ID: parsedChat.ID,
		}

		result, err := s.store.RepositoryMessages.GetChatMessages(chatEntity)

		if err != nil && err == sql.ErrNoRows {
			s.error(rw, r, http.StatusInternalServerError, fmt.Errorf("messages not found"))
			return
		}

		if err != nil {
			s.error(rw, r, http.StatusInternalServerError, fmt.Errorf("internal server error while getting user's chats"))
			return
		}

		jsonResult, err := json.Marshal(result)

		if err != nil {
			s.error(rw, r, http.StatusInternalServerError, fmt.Errorf("internal server error while parsing user's chats"))
			return
		}

		s.respond(rw, r, http.StatusOK, string(jsonResult))
	}
}
