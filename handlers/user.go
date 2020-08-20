package handlers

import (
	"github.com/google/logger"
	json "github.com/mailru/easyjson"
	"github.com/saskamegaprogrammist/amazingChat/models"
	"github.com/saskamegaprogrammist/amazingChat/useCases"
	"github.com/saskamegaprogrammist/amazingChat/utils"
	"net/http"
)

type UsersHandlers struct {
	UsersUC *useCases.UsersUC
}

func (uh *UsersHandlers) Add(writer http.ResponseWriter, req *http.Request) {
	var newUser models.User
	err := json.UnmarshalFromReader(req.Body, &newUser)
	if err != nil {
		logger.Errorf("Error unmarshaling json: %v", err)
		utils.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}
	usersExisting, err := uh.UsersUC.Add(&newUser)
	if usersExisting {
		utils.CreateErrorAnswerJson(writer, utils.StatusCode("Conflict"), models.CreateMessage(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		utils.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}

	utils.CreateAnswerIdJson(writer,  utils.StatusCode("Created"), models.CreateId(newUser.Id))
}
