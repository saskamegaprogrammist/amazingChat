package repository

import (
	"fmt"
	"github.com/google/logger"
	"github.com/saskamegaprogrammist/amazingChat/models"
	"github.com/saskamegaprogrammist/amazingChat/utils"
	"strconv"
)

type UsersRepo struct {
}

func (usersRepo *UsersRepo) GetUserIdByUsername(user *models.User) (int, error) {
	userExistsId := utils.ERROR_ID
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		logger.Errorf("Failed to start transaction: %v", err)
		return userExistsId, err
	}

	row := transaction.QueryRow("SELECT id FROM users WHERE username = $1", user.UserName)
	row.Scan(&userExistsId)

	err = transaction.Commit()
	if err != nil {
		logger.Errorf("Error commit: %v", err)
		return userExistsId, err
	}
	return userExistsId, nil
}

func (usersRepo *UsersRepo) CheckUser(userId string) (int, error) {
	userIdExists := utils.ERROR_ID

	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		logger.Errorf("Failed to start transaction: %v", err)
		return userIdExists, err
	}

	row := transaction.QueryRow("SELECT id FROM users WHERE id = $1", userId)
	err = row.Scan(&userIdExists)
	if err != nil {
		logger.Errorf("Failed to scan row: %v", err)
		errRollback := transaction.Rollback()
		if errRollback != nil {
			logger.Errorf("Failed to rollback: %v", err)
		}
		return userIdExists, fmt.Errorf("this user doesn't exist")
	}
	err = transaction.Commit()
	if err != nil {
		logger.Errorf("Error commit: %v", err)
		return userIdExists, err
	}
	return userIdExists, err
}

func (usersRepo *UsersRepo) InsertUser(user *models.User) error {
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		logger.Errorf("Failed to start transaction: %v", err)
		return err
	}
	var userId int
	row := transaction.QueryRow("INSERT INTO users (username, created) VALUES ($1, $2) RETURNING id",
		user.UserName, user.Created)
	err = row.Scan(&userId)
	if err != nil {
		logger.Errorf("Failed to scan row: %v", err)
		errRollback := transaction.Rollback()
		if errRollback != nil {
			logger.Errorf("Failed to rollback: %v", err)
		}
		return err
	}
	user.Id = strconv.Itoa(userId)
	err = transaction.Commit()
	if err != nil {
		logger.Errorf("Error commit: %v", err)
		return err
	}
	return nil
}
