package repository

import (
	"github.com/saskamegaprogrammist/amazingChat/models"
)

type MessagesRepoInterface interface {
	InsertMessage(message *models.Message) (int, error)
	GetMessagesByChatId(chatId string) ([]models.Message, error)
}

