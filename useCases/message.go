package useCases

import "github.com/saskamegaprogrammist/amazingChat/repository"

type MessageUC struct {
	MessagesRepo *repository.MessagesRepo
}