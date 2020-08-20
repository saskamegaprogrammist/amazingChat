package useCases

import "github.com/saskamegaprogrammist/amazingChat/repository"

type UseCases struct {
	UsersUC *UsersUC
	ChatsUC *ChatsUC
	MessagesUC *MessagesUC
}

var uc UseCases

func Init(usersRepo *repository.UsersRepo, chatsRepo *repository.ChatsRepo, messagesRepo *repository.MessagesRepo) error {
	uc.UsersUC = &UsersUC{usersRepo}
	uc.ChatsUC = &ChatsUC{chatsRepo, usersRepo}
	uc.MessagesUC = &MessagesUC{messagesRepo, chatsRepo, usersRepo}

	return nil
}

func GetUsersUC() *UsersUC {
	return uc.UsersUC
}

func GetChatsUC() *ChatsUC {
	return uc.ChatsUC
}
func GetMessagesUC() *MessagesUC {
	return uc.MessagesUC
}