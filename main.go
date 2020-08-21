package main

import (
	"github.com/google/logger"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
	chat_handlers "github.com/saskamegaprogrammist/amazingChat/handlers"
	"github.com/saskamegaprogrammist/amazingChat/useCases"
	"github.com/saskamegaprogrammist/amazingChat/utils"
	"github.com/saskamegaprogrammist/amazingChat/repository"
	"net/http"
	"time"
)

func main() {

	utils.LoggerSetup()
	defer utils.LoggerClose()

	err := repository.Init(pgx.ConnConfig{
		Database: utils.DBName,
		Host:     "localhost",
		User:     "alexis",
		Password: "sinope27",
	})
	if err != nil {
		logger.Fatalf("Couldn't initialize database: %v", err)
	}

	err = useCases.Init(repository.GetUsersRepo(), repository.GetChatsRepo(), repository.GetMessagesRepo())
	if err != nil {
		logger.Fatalf("Couldn't initialize useCases: %v", err)
	}

	err = chat_handlers.Init(useCases.GetUsersUC(), useCases.GetChatsUC(), useCases.GetMessagesUC())
	if err != nil {
		logger.Fatalf("Couldn't initialize handlers: %v", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/users/add",  chat_handlers.GetUsersH().Add).Methods("POST")
	r.HandleFunc("/chats/add",  chat_handlers.GetChatsH().Add).Methods("POST")
	r.HandleFunc("/messages/add",  chat_handlers.GetMessagesH().Add).Methods("POST")
	r.HandleFunc("/chats/get",  chat_handlers.GetChatsH().Get).Methods("POST")
	//r.HandleFunc("/messages/get",  chat_handlers.GetMessagesH().Get).Methods("POST")

	cors := handlers.CORS(handlers.AllowedOrigins([]string{"http://localhost:5000"}), handlers.AllowCredentials(), handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"}))

	server := &http.Server{
		Addr: utils.PortNum,
		Handler : cors(r),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err = server.ListenAndServe()

	if err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}