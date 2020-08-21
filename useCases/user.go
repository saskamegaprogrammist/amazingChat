package useCases

import (
	"fmt"
	"github.com/saskamegaprogrammist/amazingChat/models"
	"github.com/saskamegaprogrammist/amazingChat/repository"
	"github.com/saskamegaprogrammist/amazingChat/utils"
	"time"
)

type UsersUC struct {
	UsersRepo *repository.UsersRepo
}

func (userUC *UsersUC) Add(user *models.User) (bool, error) {
	id, err := userUC.UsersRepo.GetUserIdByUsername(user)
	if err != nil {
		return false, err
	}
	if id != utils.ERROR_ID {
		return true, fmt.Errorf("this username is already taken")
	} else {
		user.Created = time.Now()
		return false, userUC.UsersRepo.InsertUser(user)
	}
}
