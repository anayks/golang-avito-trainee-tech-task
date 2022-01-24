package server

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	db "github.com/anayks/golang-avito-trainee-tech-task/internal/app/store"
	store "github.com/anayks/golang-avito-trainee-tech-task/internal/app/store/repository"
)

type ctxKey int8

const (
	ctxKeyUser ctxKey = iota
	ctxKeyRequestID
)

type server struct {
	logger *logrus.Logger
	router *mux.Router
	store  store.Store
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.Use(s.accessLogMiddleware)
	s.router.Use(s.panicMiddleware)
	s.router.HandleFunc("/users/add", s.handleAddUser).Methods("POST")
	s.router.HandleFunc("/chats/add", s.handlerCreateChat).Methods("POST")
	s.router.HandleFunc("/messages/add", s.handlerSendMessage).Methods("POST")
	s.router.HandleFunc("/messages/get", s.handlerGetChatMessages()).Methods("POST")
	s.router.HandleFunc("/chats/get", s.handlerGetUserListOfChats()).Methods("POST")
}

func newServer(store *store.Store) (s *server) {
	s = &server{
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  *store,
	}
	s.configureRouter()
	return s
}

func StartServer() {

	db := db.Connect()
	defer db.Close()
	sqlstore := store.New(db)
	server := newServer(sqlstore)

	server.logger.Println("Server started at 9000 port!")
	http.ListenAndServe(":9000", server)
}
