package useCases

import (
	"fmt"
	"github.com/saskamegaprogrammist/amazingChat/models"
	"github.com/saskamegaprogrammist/amazingChat/repository"
	"github.com/saskamegaprogrammist/amazingChat/utils"
	"time"
)

type MessagesUC struct {
	MessagesRepo *repository.MessagesRepo
	ChatsRepo *repository.ChatsRepo
	UsersRepo *repository.UsersRepo
}

func (messagesUC *MessagesUC) Add(message *models.Message) (bool, error) {
	message.Created = time.Now()
	chatExistsId, err := messagesUC.ChatsRepo.CheckUserInChat(message.Author, message.Chat)
	if err != nil {
		return false, err
	}
	if chatExistsId == utils.ERROR_ID {
		return true, fmt.Errorf("this user is not in this chat")
	} else {
		errType, err := messagesUC.MessagesRepo.InsertMessage(message)
		if errType == utils.NO_ERROR {
			return false, nil
		} else if errType == utils.USER_ERROR {
			return true, err
		} else {
			return false, err
		}
	}
}
