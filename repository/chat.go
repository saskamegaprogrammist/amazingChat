package repository

import (
	"fmt"
	"github.com/google/logger"
	"github.com/jackc/pgx"
	"github.com/saskamegaprogrammist/amazingChat/models"
	"github.com/saskamegaprogrammist/amazingChat/utils"
	"strconv"
)

type ChatsRepo struct {
}

func (chatsRepo *ChatsRepo) GetChatIdByName(chat *models.Chat) (int, error) {
	chatExistsId := utils.ERROR_ID
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		logger.Errorf("Failed to start transaction: %v", err)
		return chatExistsId, err
	}

	row := transaction.QueryRow("SELECT id FROM chats WHERE name = $1", chat.Name)
	row.Scan(&chatExistsId)

	err = transaction.Commit()
	if err != nil {
		logger.Errorf("Error commit: %v", err)
		return chatExistsId, err
	}
	return chatExistsId, nil
}

func (chatsRepo *ChatsRepo) CheckUserInChat(userId string, chatId string) (int, error) {
	chatIdExists := utils.ERROR_ID

	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		logger.Errorf("Failed to start transaction: %v", err)
		return chatIdExists, err
	}

	row := transaction.QueryRow("SELECT chatid FROM chat_users WHERE chatid = $1 and userid = $2", chatId, userId)
	err = row.Scan(&chatIdExists)
	if err != nil {
		logger.Errorf("Failed to scan row: %v", err)
		errRollback := transaction.Rollback()
		if errRollback != nil {
			logger.Errorf("Failed to rollback: %v", err)
		}
		return chatIdExists, fmt.Errorf("this user is not in this chat or this chat doesn't exist")
	}
	err = transaction.Commit()
	if err != nil {
		logger.Errorf("Error commit: %v", err)
		return chatIdExists, err
	}
	return chatIdExists, err
}

func (chatsRepo *ChatsRepo) InsertChat(chat *models.Chat) (int, error) {
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		logger.Errorf("Failed to start transaction: %v", err)
		return utils.SERVER_ERROR, err
	}
	var chatId int
	row := transaction.QueryRow("INSERT INTO chats (name, created) VALUES ($1, $2) RETURNING id",
		chat.Name, chat.Created)
	err = row.Scan(&chatId)
	if err != nil {
		logger.Errorf("Failed to scan row: %v", err)
		errRollback := transaction.Rollback()
		if errRollback != nil {
			logger.Errorf("Failed to rollback: %v", err)
			return utils.SERVER_ERROR, errRollback
		}
		return utils.SERVER_ERROR, err
	}
	chat.Id = strconv.Itoa(chatId)
	errType, err := chatsRepo.insertChatUsers(chat, transaction)
	if err != nil {
		return errType, err
	}
	err = transaction.Commit()
	if err != nil {
		logger.Errorf("Error commit: %v", err)
		return utils.SERVER_ERROR, err
	}
	return utils.NO_ERROR, err
}

func (chatsRepo *ChatsRepo) insertChatUsers(chat *models.Chat, transaction *pgx.Tx) (int, error) {
	for _, userId := range chat.Users {
		_, err := transaction.Exec("INSERT INTO chat_users (chatid, userid) VALUES ($1, $2)", chat.Id, userId)
		if err != nil {
			logger.Errorf("Failed to insert to chat_users: %v", err)
			errRollback := transaction.Rollback()
			if errRollback != nil {
				logger.Errorf("Failed to rollback: %v", err)
				return utils.SERVER_ERROR, errRollback
			}
			return utils.USER_ERROR, fmt.Errorf("one of the users doesn't exist or chat couldn't be created")
		}
	}
	return utils.NO_ERROR, nil
}

func (chatsRepo *ChatsRepo) GetChatsByUserId(userId string) ([]models.Chat, error) {
	chats := make([]models.Chat, 0)
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		logger.Errorf("Failed to start transaction: %v", err)
		return chats, err
	}

	queryString := `SELECT id, name, created, users FROM (SELECT chatid, MAX(created) as max_created FROM messages GROUP BY chatid) as mess RIGHT OUTER JOIN
	(SELECT chatid FROM chat_users WHERE userid=$1) as user_chats on mess.chatid = user_chats.chatid JOIN chats on user_chats.chatid = chats.id JOIN (SELECT array_agg(userid::text) as users, chatid FROM chat_users GROUP BY chatid) as r
	on r.chatid = user_chats.chatid ORDER BY coalesce(max_created, created) DESC`
	rows, err := transaction.Query(queryString, userId)
	if err != nil {
		return chats, nil
	}
	for rows.Next() {
		var cFound models.Chat
		var chatId int
		err = rows.Scan(&chatId, &cFound.Name, &cFound.Created, &cFound.Users)
		if err != nil {
			logger.Errorf("Failed to retrieve chat: %v", err)
			errRollback := transaction.Rollback()
			if errRollback != nil {
				logger.Errorf("Failed to rollback: %v", err)
				return chats, errRollback
			}
			return chats, err
		}
		cFound.Id = strconv.Itoa(chatId)
		chats = append(chats, cFound)
	}
	err = transaction.Commit()
	if err != nil {
		logger.Errorf("Error commit: %v", err)
		return chats, err
	}
	return chats, nil
}

func (chatsRepo *ChatsRepo) CheckChat(chatId string) (int, error) {
	chatIdExists := utils.ERROR_ID

	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		logger.Errorf("Failed to start transaction: %v", err)
		return chatIdExists, err
	}

	row := transaction.QueryRow("SELECT id FROM chats WHERE id = $1", chatId)
	err = row.Scan(&chatIdExists)
	if err != nil {
		logger.Errorf("Failed to scan row: %v", err)
		errRollback := transaction.Rollback()
		if errRollback != nil {
			logger.Errorf("Failed to rollback: %v", err)
		}
		return chatIdExists, fmt.Errorf("this chat doesn't exist")
	}
	err = transaction.Commit()
	if err != nil {
		logger.Errorf("Error commit: %v", err)
		return chatIdExists, err
	}
	return chatIdExists, err
}
