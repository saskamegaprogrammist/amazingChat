package handlers

import "github.com/saskamegaprogrammist/amazingChat/useCases"

type Handlers struct {
	UsersHandlers *UsersHandlers
	ChatsHandlers *ChatsHandlers
	MessagesHandlers *MessagesHandlers
}
var h Handlers

func Init(usersUC *useCases.UsersUC, chatsUC *useCases.ChatsUC, messagesUC *useCases.MessagesUC) error {
	h.UsersHandlers = &UsersHandlers{usersUC}
	h.ChatsHandlers = &ChatsHandlers{chatsUC}
	h.MessagesHandlers = &MessagesHandlers{messagesUC}
	return nil
}

func GetUsersH() *UsersHandlers {
	return h.UsersHandlers
}

func GetChatsH() *ChatsHandlers {
	return h.ChatsHandlers
}

func GetMessagesH() *MessagesHandlers {
	return h.MessagesHandlers
}

