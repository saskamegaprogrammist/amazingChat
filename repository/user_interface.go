package repository

import "github.com/saskamegaprogrammist/amazingChat/models"

type UsersRepoInterface interface {
	GetUserIdByUsername(user *models.User) (int, error)
	CheckUser(userId string) (int, error)
	InsertUser(user *models.User) error
}
