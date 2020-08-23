package useCases

import "github.com/saskamegaprogrammist/amazingChat/models"

type MessagesUCInterface interface {
	Add(message *models.Message) (bool, error)
	GetChatMessagesSorted(chat *models.ChatId) (bool, []models.Message, error)
}
