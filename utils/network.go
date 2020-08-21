package utils
import (
	"github.com/google/logger"
	json "github.com/mailru/easyjson"
	chat_models "github.com/saskamegaprogrammist/amazingChat/models"
	"net/http"
)

func createAnswerJson(w http.ResponseWriter, statusCode int, data []byte)  {
	w.WriteHeader(statusCode)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")
	_, err := w.Write(data)
	if err != nil {
		logger.Errorf( "Error writing answer: %v", err)
	}
}

func CreateErrorAnswerJson(writer http.ResponseWriter, statusCode int, error chat_models.RequestError) {
	marshalledError, err := json.Marshal(error)
	if err != nil {
		logger.Errorf( "Error marhalling json: %v", err)
	}
	createAnswerJson(writer, statusCode, marshalledError)
}

func CreateAnswerIdJson(writer http.ResponseWriter, statusCode int, id chat_models.IdModel) {
	marshalledId, err := json.Marshal(id)
	if err != nil {
		logger.Errorf( "Error marhalling json: %v", err)
	}
	createAnswerJson(writer, statusCode, marshalledId)
}

func CreateAnswerChatsJson(writer http.ResponseWriter, statusCode int, chats chat_models.Chats) {
	marshalledChats, err := json.Marshal(chats)
	if err != nil {
		logger.Errorf( "Error marhalling json: %v", err)
	}
	createAnswerJson(writer, statusCode, marshalledChats)
}

func CreateAnswerMessagesJson(writer http.ResponseWriter, statusCode int, messages chat_models.Messages) {
	marshalledMessages, err := json.Marshal(messages)
	if err != nil {
		logger.Errorf( "Error marhalling json: %v", err)
	}
	createAnswerJson(writer, statusCode, marshalledMessages)
}
