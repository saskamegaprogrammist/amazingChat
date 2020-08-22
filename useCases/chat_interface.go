package useCases

import "github.com/saskamegaprogrammist/amazingChat/models"

type ChatsUCInterface interface {
	Add(chat *models.Chat) (bool, bool, error)
	GetUserChatsSorted(user *models.UserId) (bool, []models.Chat, error)
}
