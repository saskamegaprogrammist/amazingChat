package repository

import (
	"fmt"
	"github.com/google/logger"
	"github.com/saskamegaprogrammist/amazingChat/models"
	"github.com/saskamegaprogrammist/amazingChat/utils"
	"strconv"
)

type MessagesRepo struct {

}

func (messagesRepo *MessagesRepo) InsertMessage(message *models.Message) (int, error) {
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		logger.Errorf("Failed to start transaction: %v", err)
		return utils.SERVER_ERROR, err
	}

	var messageId int
	row := transaction.QueryRow("INSERT INTO messages (text, chatid, userid, created) VALUES ($1, $2, $3, $4) RETURNING id",
		message.Text, message.Chat, message.Author, message.Created)
	err = row.Scan(&messageId)
	if err != nil {
		logger.Errorf("Failed to scan row: %v", err)
		errRollback := transaction.Rollback()
		if errRollback != nil {
			logger.Errorf("Failed to rollback: %v", err)
			return utils.SERVER_ERROR, errRollback
		}
		return utils.USER_ERROR, fmt.Errorf("this chat or user doesn't exist")
	}
	message.Id = strconv.Itoa(messageId)

	err = transaction.Commit()
	if err != nil {
		logger.Errorf("Error commit: %v", err)
		return utils.SERVER_ERROR, err
	}
	return utils.NO_ERROR, nil
}

