package auto

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	ChatEntity "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/chat"
	ChatMessage "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/message"
	user "github.com/anayks/golang-avito-trainee-tech-task/internal/app/entity/user"
)

func (s *server) handleAddUser(rw http.ResponseWriter, r *http.Request) {
	parsedUser := &user.ChatUser{}

	if err := json.NewDecoder(r.Body).Decode(parsedUser); err != nil {
		s.error(rw, r, http.StatusBadRequest, err)
		return
	}

	matched, _ := regexp.Match(`^[a-zA-Z0-9а-яА-Я_]{4,20}$`, []byte(parsedUser.Username))

	if !matched {
		s.error(rw, r, http.StatusUnprocessableEntity, fmt.Errorf("invalid username characters. Actually: %s", parsedUser.Username))
		return
	}

	id := s.store.Users().Create(parsedUser)

	fmt.Fprint(rw, id)
}

func (s *server) handlerCreateChat(rw http.ResponseWriter, r *http.Request) {
	parsedChat := &ChatEntity.Chat{}

	if err := json.NewDecoder(r.Body).Decode(&parsedChat); err != nil {
		s.error(rw, r, http.StatusUnprocessableEntity, err)
		return
	}

	matched, _ := regexp.Match(`^[a-zA-Z0-9а-яА-Я_]{4,20}$`, []byte(parsedChat.Name))

	if !matched {
		s.error(rw, r, http.StatusUnprocessableEntity, fmt.Errorf("invalid chat name characters. Actually: %s", parsedChat.Name))
		return
	}

	id, err := s.store.Chats().Create(r.Context(), parsedChat)

	if err != nil {
		s.ErrorLog(r, err)
		s.error(rw, r, http.StatusUnprocessableEntity, fmt.Errorf("internal error while creating chat"))
		return
	}

	fmt.Fprint(rw, id)
}

func (s *server) handlerSendMessage(rw http.ResponseWriter, r *http.Request) {
	parsedMessage := &ChatMessage.Message{}

	if err := json.NewDecoder(r.Body).Decode(&parsedMessage); err != nil {
		s.error(rw, r, http.StatusUnprocessableEntity, err)
		return
	}

	if parsedMessage.Chat <= 0 {
		s.error(rw, r, http.StatusUnprocessableEntity, fmt.Errorf("chatID is not valid. Actually: %v", parsedMessage.Chat))
		return
	}

	if parsedMessage.Author <= 0 {
		s.error(rw, r, http.StatusUnprocessableEntity, fmt.Errorf("author is not valid. Actually: %v", parsedMessage.Author))
		return
	}

	if len(parsedMessage.Text) == 0 {
		s.error(rw, r, http.StatusUnprocessableEntity, fmt.Errorf("message text cannot be empty"))
		return
	}

	id, err := s.store.Messages().Create(r.Context(), parsedMessage)

	if err != nil {
		s.ErrorLog(r, err)
		s.error(rw, r, http.StatusUnprocessableEntity, fmt.Errorf("internal server error while creating message"))
		return
	}

	fmt.Fprint(rw, id)
}

func (s *server) handlerGetUserListOfChats(rw http.ResponseWriter, r *http.Request) {

	parsedUser := &user.ChatUser{}

	if err := json.NewDecoder(r.Body).Decode(&parsedUser); err != nil {
		s.error(rw, r, http.StatusBadRequest, err)
		return
	}

	if parsedUser.ID < 0 {
		s.error(rw, r, http.StatusUnprocessableEntity, fmt.Errorf("user is not valid. Actually: %v", parsedUser.ID))
		return
	}

	result, err := s.store.Chats().GetUserChats(parsedUser)

	if err != nil {
		s.ErrorLog(r, err)
		s.error(rw, r, http.StatusUnprocessableEntity, fmt.Errorf("internal server error while getting user's chats"))
		return
	}

	fmt.Fprint(rw, result)
}

func (s *server) handlerGetChatMessages(rw http.ResponseWriter, r *http.Request) {
	type parseServer struct {
		ID int64 `json:"chat"`
	}

	parsedChat := &parseServer{}

	if err := json.NewDecoder(r.Body).Decode(&parsedChat); err != nil {
		s.error(rw, r, http.StatusBadRequest, err)
		return
	}

	if parsedChat.ID <= 0 {
		s.error(rw, r, http.StatusUnprocessableEntity, fmt.Errorf("chat-id is not valid. Actually: %v", parsedChat.ID))
		return
	}

	chatEntity := &ChatEntity.Chat{
		ID: parsedChat.ID,
	}

	result, err := s.store.Messages().GetChatMessages(chatEntity)

	if err != nil {
		if err == sql.ErrNoRows {
			s.ErrorLog(r, err)
			s.error(rw, r, http.StatusNotFound, fmt.Errorf("messages not found"))
			return
		}

		s.ErrorLog(r, err)
		s.error(rw, r, http.StatusNotFound, fmt.Errorf("internal server error while getting user's chats"))
		return
	}

	fmt.Fprint(rw, result)
}
