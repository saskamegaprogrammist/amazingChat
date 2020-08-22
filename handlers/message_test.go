package handlers

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/jackc/fake"
	"github.com/saskamegaprogrammist/amazingChat/models"
	"github.com/saskamegaprogrammist/amazingChat/useCases"
	"github.com/saskamegaprogrammist/amazingChat/utils"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"net/http"
	"testing"
	"time"
)

var mh MessagesHandlers

var testMessageOne = models.Message{
	Chat: "3",
	Author: "2",
	Text: fake.Sentence(),
}

var testMessageWrong = models.Message{
	Chat: "5",
	Author: "2",
	Text: fake.Sentence(),
}

var testChatIdOne = models.ChatId{
	ChatId:"1",
}

var testChatIdWrong = models.ChatId{
	ChatId:"15",
}

var testChatOneMessages =  []models.Message{
	{Id: "2", Chat: "1", Author: "1", Text: "sss", Created : time.Now()},
	{Id: "5", Chat: "1", Author: "2", Text: "sdfeeei", Created : time.Now()},
	{Id: "6", Chat: "1", Author: "2", Text: "sdfeeei", Created : time.Now()},
	{Id: "7", Chat: "1", Author: "2", Text: "sdfeeei", Created : time.Now()},
}

func TestAddMessage(t *testing.T) {
	t.Run("MessageAddOK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := useCases.NewMockMessagesUCInterface(ctrl)
		mockUseCase.EXPECT().Add(&testMessageOne).Return(false, nil)

		mh.MessagesUC = mockUseCase

		jsonBody := fmt.Sprintf(`{"chat": "%s", "author": "%s", "text": "%s"}`, testMessageOne.Chat, testMessageOne.Author, testMessageOne.Text)

		apitest.New("MessageAddOK").
			Handler(http.HandlerFunc(mh.Add)).
			Method("Post").
			URL(utils.GetAPIAddress("addMessage")).
			Body(jsonBody).
			Expect(t).
			Status(http.StatusCreated).
			Assert(jsonpath.Matches(`$.id`, `^[0-9]*$`)).
			End()
	})

	t.Run("UserError", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := useCases.NewMockMessagesUCInterface(ctrl)
		mockUseCase.EXPECT().Add(&testMessageWrong).Return(true, errors.New("some user error"))

		mh.MessagesUC = mockUseCase

		jsonBody := fmt.Sprintf(`{"chat": "%s", "author": "%s", "text": "%s"}`, testMessageWrong.Chat, testMessageWrong.Author, testMessageWrong.Text)

		apitest.New("UserError").
			Handler(http.HandlerFunc(mh.Add)).
			Method("Post").
			URL(utils.GetAPIAddress("addMessage")).
			Body(jsonBody).
			Expect(t).
			Status(http.StatusBadRequest).
			Assert(jsonpath.Present(`$.message`)).
			End()
	})


	t.Run("InternalError", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := useCases.NewMockMessagesUCInterface(ctrl)
		mockUseCase.EXPECT().Add(&testMessageWrong).Return(false, errors.New("database error"))

		mh.MessagesUC = mockUseCase

		jsonBody := fmt.Sprintf(`{"chat": "%s", "author": "%s", "text": "%s"}`, testMessageWrong.Chat, testMessageWrong.Author, testMessageWrong.Text)

		apitest.New("UserError").
			Handler(http.HandlerFunc(mh.Add)).
			Method("Post").
			URL(utils.GetAPIAddress("addMessage")).
			Body(jsonBody).
			Expect(t).
			Status(http.StatusInternalServerError).
			Assert(jsonpath.Present(`$.message`)).
			End()

	})

	t.Run("MalformedJSON", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		apitest.New("MalformedJSON").
			Handler(http.HandlerFunc(mh.Add)).
			Method("Post").
			URL(utils.GetAPIAddress("addMessage")).
			Body(`{"chat" "author": text": "}`).
			Expect(t).
			Status(http.StatusInternalServerError).
			Assert(jsonpath.Contains(`$.message`, "Error unmarshaling json")).
			End()

	})
}

func TestGetChatMessagesSorted(t *testing.T) {
	t.Run("MessagesGetOK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := useCases.NewMockMessagesUCInterface(ctrl)
		mockUseCase.EXPECT().GetChatMessagesSorted(&testChatIdOne).Return(false, testChatOneMessages, nil)

		mh.MessagesUC = mockUseCase

		jsonBody := fmt.Sprintf(`{"chat": "%s"}`, testChatIdOne.ChatId)

		apitest.New("MessagesGetOK").
			Handler(http.HandlerFunc(mh.Get)).
			Method("Post").
			URL(utils.GetAPIAddress("getMessages")).
			Body(jsonBody).
			Expect(t).
			Status(http.StatusOK).
			End()
	})

	t.Run("ChatNotExist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := useCases.NewMockMessagesUCInterface(ctrl)
		mockUseCase.EXPECT().GetChatMessagesSorted(&testChatIdWrong).Return(true, []models.Message{}, errors.New("chat doesn't exist"))

		mh.MessagesUC = mockUseCase

		jsonBody := fmt.Sprintf(`{"chat": "%s"}`, testChatIdWrong.ChatId)

		apitest.New("ChatNotExist").
			Handler(http.HandlerFunc(mh.Get)).
			Method("Post").
			URL(utils.GetAPIAddress("getMessages")).
			Body(jsonBody).
			Expect(t).
			Status(http.StatusBadRequest).
			Assert(jsonpath.Present(`$.message`)).
			End()
	})

	t.Run("InternalError", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := useCases.NewMockMessagesUCInterface(ctrl)
		mockUseCase.EXPECT().GetChatMessagesSorted(&testChatIdOne).Return(false, []models.Message{}, errors.New("database error"))

		mh.MessagesUC = mockUseCase

		jsonBody := fmt.Sprintf(`{"chat": "%s"}`, testChatIdOne.ChatId)

		apitest.New("InternalError").
			Handler(http.HandlerFunc(mh.Get)).
			Method("Post").
			URL(utils.GetAPIAddress("getMessages")).
			Body(jsonBody).
			Expect(t).
			Status(http.StatusInternalServerError).
			Assert(jsonpath.Present(`$.message`)).
			End()
	})

	t.Run("MalformedJSON", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		apitest.New("MalformedJSON").
			Handler(http.HandlerFunc(mh.Get)).
			Method("Post").
			URL(utils.GetAPIAddress("getMessages")).
			Body(`{"chat": "}`).
			Expect(t).
			Status(http.StatusInternalServerError).
			Assert(jsonpath.Contains(`$.message`, "Error unmarshaling json")).
			End()

	})

}
