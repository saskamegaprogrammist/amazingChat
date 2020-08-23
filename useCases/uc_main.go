package useCases

import "github.com/saskamegaprogrammist/amazingChat/repository"

type UseCases struct {
	UsersUC    *UsersUC
	ChatsUC    *ChatsUC
	MessagesUC *MessagesUC
}

var uc UseCases

func Init(usersRepo repository.UsersRepoInterface, chatsRepo repository.ChatsRepoInterface, messagesRepo repository.MessagesRepoInterface) error {
	uc.UsersUC = &UsersUC{usersRepo}
	uc.ChatsUC = &ChatsUC{chatsRepo, usersRepo}
	uc.MessagesUC = &MessagesUC{messagesRepo, chatsRepo, usersRepo}

	return nil
}

func GetUsersUC() UsersUCInterface {
	return uc.UsersUC
}

func GetChatsUC() ChatsUCInterface {
	return uc.ChatsUC
}
func GetMessagesUC() MessagesUCInterface {
	return uc.MessagesUC
}
