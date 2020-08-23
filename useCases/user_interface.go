package useCases

import "github.com/saskamegaprogrammist/amazingChat/models"

type UsersUCInterface interface {
	Add(user *models.User) (bool, error)
}
