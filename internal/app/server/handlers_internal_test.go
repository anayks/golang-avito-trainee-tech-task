package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anayks/golang-avito-trainee-tech-task/internal/app/server"
	sqlstore "github.com/anayks/golang-avito-trainee-tech-task/internal/app/store/repository"
	"github.com/stretchr/testify/assert"
)

func TestServer_handleAddUser(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, "../../../.env")
	defer teardown("users")
	sqlstore := sqlstore.New(db)
	s := server.NewServer(sqlstore)

	testData := []struct {
		Name         string
		Data         interface{}
		expectedCode int
	}{
		{
			Name: "addUser 0",
			Data: map[string]string{
				"username": "user_1",
			},
			expectedCode: http.StatusOK,
		},
		{
			Name: "addUser 1",
			Data: map[string]string{
				"username": "test135346хей",
			},
			expectedCode: http.StatusOK,
		},
		{
			Name:         "addUser 2",
			Data:         nil,
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, val := range testData {
		t.Run(val.Name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(val.Data)
			req, err := http.NewRequest("POST", "/users/add", b)
			s.ServeHTTP(rec, req)
			assert.NoError(t, err)
			assert.Equal(t, val.expectedCode, rec.Code)
		})
	}
}

func TestServer_handlerCreateChat(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, "../../../.env")
	defer teardown("chats")
	defer teardown("chatsUsers")
	sql := sqlstore.New(db)
	s := server.NewServer(sql)
	_, user, _, teardownFiller := sqlstore.DBFiller(t, db, teardown)
	defer teardownFiller()

	testData := []struct {
		Name         string
		Data         interface{}
		expectedCode int
	}{
		{
			Name: "createChat 0",
			Data: map[string]interface{}{
				"name":  "chat_1",
				"users": []int64{user.ID},
			},
			expectedCode: http.StatusOK,
		},
		{
			Name: "createChat 1",
			Data: map[string]interface{}{
				"name":  "34@$@#%#@^$#^%$^#$^$%*^%^&(*&^(&^)&*)&)^%$#%#$%",
				"users": []int64{-1, -2},
			},
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, val := range testData {
		t.Run(val.Name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(val.Data)
			req, err := http.NewRequest("POST", "/chats/add", b)
			s.ServeHTTP(rec, req)
			assert.NoError(t, err)
			assert.Equal(t, val.expectedCode, rec.Code)
		})
	}
}

func TestServer_handlerSendMessage(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, "../../../.env")
	defer teardown("messages")
	sql := sqlstore.New(db)
	s := server.NewServer(sql)
	chat_id, user, _, teardownFiller := sqlstore.DBFiller(t, db, teardown)
	defer teardownFiller()

	var longGeneratedText string

	for i := 0; i < 320; i++ {
		longGeneratedText = fmt.Sprintf("%s %s", longGeneratedText, "test")
	}

	testData := []struct {
		Name         string
		UserID       int64
		Data         interface{}
		expectedCode int
	}{
		{
			Name: "sendMessage 0",
			Data: map[string]interface{}{
				"author": user.ID,
				"chat":   chat_id,
				"text":   "Heyo, buddy!",
			},
			expectedCode: http.StatusOK,
		},
		{
			Name: "sendMessage 1",
			Data: map[string]interface{}{
				"author": user.ID,
				"chat":   chat_id,
				"text":   longGeneratedText,
			},
			expectedCode: http.StatusInternalServerError,
		},
		{
			Name: "sendMessage 2",
			Data: map[string]interface{}{
				"author": user.ID + 1,
				"chat":   chat_id,
				"text":   longGeneratedText,
			},
			expectedCode: http.StatusInternalServerError,
		},
		{
			Name:         "sendMessage 3",
			Data:         "",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, val := range testData {
		t.Run(val.Name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(val.Data)
			req, err := http.NewRequest("POST", "/messages/add", b)
			s.ServeHTTP(rec, req)
			assert.NoError(t, err)
			assert.Equal(t, val.expectedCode, rec.Code)
		})
	}
}

func TestServer_GetUserChats(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, "../../../.env")
	defer teardown("messages")
	sql := sqlstore.New(db)
	s := server.NewServer(sql)
	_, user, _, teardownFiller := sqlstore.DBFiller(t, db, teardown)
	defer teardownFiller()

	testData := []struct {
		Name         string
		UserID       int64
		Data         interface{}
		expectedCode int
	}{
		{
			Name: "getMessage 0",
			Data: map[string]interface{}{
				"user": user.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			Name: "getMessage 1",
			Data: map[string]interface{}{
				"user": user.ID + 1,
			},
			expectedCode: http.StatusOK,
		},
	}

	for _, val := range testData {
		t.Run(val.Name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(val.Data)
			req, err := http.NewRequest("POST", "/chats/get", b)
			s.ServeHTTP(rec, req)
			assert.NoError(t, err)
			assert.Equal(t, val.expectedCode, rec.Code)
		})
	}
}

func TestServer_handlerGetUserListOfChats(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, "../../../.env")
	defer teardown("messages")
	sql := sqlstore.New(db)
	s := server.NewServer(sql)
	chat_id, _, _, teardownFiller := sqlstore.DBFiller(t, db, teardown)
	defer teardownFiller()

	testData := []struct {
		Name         string
		UserID       int64
		Data         interface{}
		expectedCode int
	}{
		{
			Name: "getUserList 0",
			Data: map[string]interface{}{
				"chat": chat_id,
			},
			expectedCode: http.StatusOK,
		},
		{
			Name: "getUserList 1",
			Data: map[string]interface{}{
				"chat": chat_id + 1,
			},
			expectedCode: http.StatusOK,
		},
	}

	for _, val := range testData {
		t.Run(val.Name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(val.Data)
			req, err := http.NewRequest("POST", "/messages/get", b)
			s.ServeHTTP(rec, req)
			assert.NoError(t, err)
			assert.Equal(t, val.expectedCode, rec.Code)
		})
	}
}
