package repository

import (
	"github.com/jackc/pgx"
	"github.com/saskamegaprogrammist/amazingChat/models"
)

type ChatsRepoInterface interface {
	GetChatIdByName(chat *models.Chat) (int, error)
	CheckUserInChat(userId string, chatId string) (int, error)
	InsertChat(chat *models.Chat) (int, error)
	insertChatUsers(chat *models.Chat, transaction *pgx.Tx) (int, error)
	GetChatsByUserId(userId string) ([]models.Chat, error)
	CheckChat(chatId string) (int, error)
}
