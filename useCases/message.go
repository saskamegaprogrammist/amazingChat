package useCases

import "github.com/saskamegaprogrammist/amazingChat/repository"

type MessagesUC struct {
	MessagesRepo *repository.MessagesRepo
	ChatsRepo *repository.ChatsRepo
	UsersRepo *repository.UsersRepo
}