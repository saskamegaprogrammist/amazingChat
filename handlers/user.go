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
	UsersUC *useCases.UserUC
}

func (uh *UsersHandlers) Add(writer http.ResponseWriter, req *http.Request) {

	var newUser models.User
	err := json.UnmarshalFromReader(req.Body, &newUser)
	if err != nil {
		utils.WriteError(false, "Error unmarshaling json", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}
	usersExisting, err := uh.UsersUC.SignUp(&newUser)
	if usersExisting {
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Conflict"), models.CreateMessage(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}

	network.CreateAnswerUserJson(writer,  utils.StatusCode("Created"), newUser)
}
