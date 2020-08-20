package useCases

import "github.com/saskamegaprogrammist/amazingChat/repository"

type ChatsUC struct {
	ChatsRepo *repository.ChatsRepo
	UsersRepo *repository.UsersRepo
}
