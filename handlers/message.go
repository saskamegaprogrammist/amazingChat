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

type MessagesHandlers struct {
	MessagesUC useCases.MessagesUCInterface
}

func (mh *MessagesHandlers) Add(writer http.ResponseWriter, req *http.Request) {
	var newMessage models.Message
	err := json.UnmarshalFromReader(req.Body, &newMessage)
	if err != nil {
		jsonError := fmt.Sprintf("Error unmarshaling json: %v", err.Error())
		logger.Errorf(jsonError)
		utils.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(jsonError))
		return
	}
	userError, err := mh.MessagesUC.Add(&newMessage)
	if userError {
		logger.Error(err)
		utils.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		utils.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}

	utils.CreateAnswerIdJson(writer, utils.StatusCode("Created"), models.CreateId(newMessage.Id))
}

func (mh *MessagesHandlers) Get(writer http.ResponseWriter, req *http.Request) {
	var chat models.ChatId
	err := json.UnmarshalFromReader(req.Body, &chat)
	if err != nil {
		jsonError := fmt.Sprintf("Error unmarshaling json: %v", err.Error())
		logger.Errorf(jsonError)
		utils.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(jsonError))
		return
	}
	chatNotExist, messages, err := mh.MessagesUC.GetChatMessagesSorted(&chat)
	if chatNotExist {
		utils.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		utils.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}

	utils.CreateAnswerMessagesJson(writer, utils.StatusCode("OK"), messages)
}
