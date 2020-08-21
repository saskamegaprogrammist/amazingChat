package useCases

import (
	"fmt"
	"github.com/saskamegaprogrammist/amazingChat/models"
	"github.com/saskamegaprogrammist/amazingChat/repository"
	"github.com/saskamegaprogrammist/amazingChat/utils"
	"time"
)

type ChatsUC struct {
	ChatsRepo *repository.ChatsRepo
	UsersRepo *repository.UsersRepo
}

func (chatUC *ChatsUC) Add(chat *models.Chat) (bool, bool, error) {
	id, err := chatUC.ChatsRepo.GetChatIdByName(chat)
	if err != nil {
		return false, false, err
	}
	if id != utils.ERROR_ID {
		return true, false, fmt.Errorf("this name is already taken")
	} else {
		chat.Created = time.Now()
		errType, err := chatUC.ChatsRepo.InsertChat(chat)
		if errType == utils.NO_ERROR {
			return false, false, nil
		} else if errType == utils.USER_ERROR {
			return false, true, err
		} else {
			return false, false, err
		}
	}
}

func (chatUC *ChatsUC) GetUserChatsSorted(user *models.UserId) (bool, []models.Chat, error) {
	chats := make([]models.Chat, 0)
	userExistsId, err := chatUC.UsersRepo.CheckUser(user.UserId)
	if err != nil {
		return false, chats, err
	}
	if userExistsId == utils.ERROR_ID {
		return true, chats, fmt.Errorf("this user doesn't exist")
	} else {
		chats, err = chatUC.ChatsRepo.GetChatsByUserId(user.UserId)
		return false, chats, err
	}
}