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

	// logger initialisation
	utils.LoggerSetup()
	defer utils.LoggerClose()

	// database initialisation
	err := repository.Init(pgx.ConnConfig{
		Database: utils.DBName,
		Host:     "localhost",
		User:     "docker",
		Password: "docker",
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

	// router initialisation

	r := mux.NewRouter()
	r.HandleFunc(utils.GetAPIAddress("addUser"),  chat_handlers.GetUsersH().Add).Methods("POST")
	r.HandleFunc(utils.GetAPIAddress("addChat"), chat_handlers.GetChatsH().Add).Methods("POST")
	r.HandleFunc(utils.GetAPIAddress("addMessage"),  chat_handlers.GetMessagesH().Add).Methods("POST")
	r.HandleFunc(utils.GetAPIAddress("getChats"), chat_handlers.GetChatsH().Get).Methods("POST")
	r.HandleFunc(utils.GetAPIAddress("getMessages"),  chat_handlers.GetMessagesH().Get).Methods("POST")

	cors := handlers.CORS(handlers.AllowCredentials(), handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"}))

	// server initialisation

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