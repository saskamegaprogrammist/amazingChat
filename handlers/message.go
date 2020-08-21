package handlers

import (
	"github.com/google/logger"
	json "github.com/mailru/easyjson"
	"github.com/saskamegaprogrammist/amazingChat/models"
	"github.com/saskamegaprogrammist/amazingChat/useCases"
	"github.com/saskamegaprogrammist/amazingChat/utils"
	"net/http"
)

type MessagesHandlers struct {
	MessagesUC *useCases.MessagesUC
}

func (mh *MessagesHandlers) Add(writer http.ResponseWriter, req *http.Request) {
	var newMessage models.Message
	err := json.UnmarshalFromReader(req.Body, &newMessage)
	if err != nil {
		logger.Errorf("Error unmarshaling json: %v", err)
		utils.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
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

	utils.CreateAnswerIdJson(writer,  utils.StatusCode("OK"), models.CreateId(newMessage.Id))
}