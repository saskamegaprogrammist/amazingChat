package handlers

import (
	"fmt"
	"github.com/google/logger"
	json "github.com/mailru/easyjson"
	"github.com/saskamegaprogrammist/amazingChat/models"
	"github.com/saskamegaprogrammist/amazingChat/useCases"
	"github.com/saskamegaprogrammist/amazingChat/utils"
	"net/http"
)

type ChatsHandlers struct {
	ChatsUC useCases.ChatsUCInterface
}

func (ch *ChatsHandlers) Add(writer http.ResponseWriter, req *http.Request) {
	var newChat models.Chat
	err := json.UnmarshalFromReader(req.Body, &newChat)
	if err != nil {
		jsonError := fmt.Sprintf("Error unmarshaling json: %v", err.Error())
		logger.Errorf(jsonError)
		utils.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(jsonError))
		return
	}
	chatExisting, userError, err := ch.ChatsUC.Add(&newChat)
	if chatExisting {
		utils.CreateErrorAnswerJson(writer, utils.StatusCode("Conflict"), models.CreateMessage(err.Error()))
		return
	}
	if userError {
		utils.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		utils.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}

	utils.CreateAnswerIdJson(writer, utils.StatusCode("Created"), models.CreateId(newChat.Id))
}

func (ch *ChatsHandlers) Get(writer http.ResponseWriter, req *http.Request) {
	var user models.UserId
	err := json.UnmarshalFromReader(req.Body, &user)
	if err != nil {
		jsonError := fmt.Sprintf("Error unmarshaling json: %v", err.Error())
		logger.Errorf(jsonError)
		utils.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(jsonError))
		return
	}
	userNotExist, chats, err := ch.ChatsUC.GetUserChatsSorted(&user)
	if userNotExist {
		utils.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		utils.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}

	utils.CreateAnswerChatsJson(writer, utils.StatusCode("OK"), chats)
}
