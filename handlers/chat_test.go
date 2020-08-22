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
	"strings"
	"testing"
	"time"
)

var ch ChatsHandlers

var testChatOne = models.Chat{
	Name: fake.ProductName(),
	Users:  []string{"1", "2"},
}

var testChatTwo = models.Chat{
	Name: fake.ProductName(),
	Users:  []string{"1", "2", "3"},
}

var testChatWrong = models.Chat{
	Name: fake.ProductName(),
	Users:  []string{"1", "15"},
}

var testUserIdOne = models.UserId{
	UserId:"1",
}

var testUserIdWrong = models.UserId{
	UserId:"15",
}

var testUserOneChats =  []models.Chat{
{Id: "1", Name: "chat_1", Users :[]string{"1", "2", "3"}, Created : time.Now()},
	{Id: "4", Name: "chat_3", Users : []string{"1", "2"}, Created : time.Now()},
	{Id: "3", Name: "chat_2", Users : []string{"1", "2", "3"}, Created : time.Now()},
}

var testUserOneChatsJSON = `[{"id":"1","name":"chat_1","users":["1","2","3"],"created_at":"2020-08-22T00:10:09.887457+03:00"},
	{"id":"4","name":"chat_3","users":["1","2"],"created_at":"2020-08-22T00:11:09.879923+03:00"},
{"id":"3","name":"chat_2","users":["1","2","3"],"created_at":"2020-08-22T00:10:57.801754+03:00"}]`

func TestAddChat(t *testing.T) {
	t.Run("ChatAddOK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := useCases.NewMockChatsUCInterface(ctrl)
		mockUseCase.EXPECT().Add(&testChatOne).Return(false, false, nil)

		ch.ChatsUC = mockUseCase

		jsonBody := fmt.Sprintf(`{"name": "%s", "users": ["`, testChatOne.Name) + strings.Join(testChatOne.Users, `", "`) + `"]}`

		apitest.New("ChatAddOK").
			Handler(http.HandlerFunc(ch.Add)).
			Method("Post").
			URL(utils.GetAPIAddress("addChat")).
			Body(jsonBody).
			Expect(t).
			Status(http.StatusCreated).
			Assert(jsonpath.Matches(`$.id`, `^[0-9]*$`)).
			End()
	})

	t.Run("ChatAddExists", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := useCases.NewMockChatsUCInterface(ctrl)
		mockUseCase.EXPECT().Add(&testChatOne).Return(true, false, errors.New("already exists chat error"))

		ch.ChatsUC = mockUseCase

		jsonBody := fmt.Sprintf(`{"name": "%s", "users": ["`, testChatOne.Name) + strings.Join(testChatOne.Users, `", "`) + `"]}`

		apitest.New("ChatAddExists").
			Handler(http.HandlerFunc(ch.Add)).
			Method("Post").
			URL(utils.GetAPIAddress("addChat")).
			Body(jsonBody).
			Expect(t).
			Status(http.StatusConflict).
			End()
	})

	t.Run("UserError", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := useCases.NewMockChatsUCInterface(ctrl)
		mockUseCase.EXPECT().Add(&testChatWrong).Return(false, true, errors.New("user doesn't exist error"))

		ch.ChatsUC = mockUseCase

		jsonBody := fmt.Sprintf(`{"name": "%s", "users": ["`, testChatWrong.Name) + strings.Join(testChatWrong.Users, `", "`) + `"]}`

		apitest.New("UserError").
			Handler(http.HandlerFunc(ch.Add)).
			Method("Post").
			URL(utils.GetAPIAddress("addChat")).
			Body(jsonBody).
			Expect(t).
			Status(http.StatusBadRequest).
			End()
	})

	t.Run("InternalError", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := useCases.NewMockChatsUCInterface(ctrl)
		mockUseCase.EXPECT().Add(&testChatTwo).Return(false, false, errors.New("database error"))

		ch.ChatsUC = mockUseCase

		jsonBody := fmt.Sprintf(`{"name": "%s", "users": ["`, testChatTwo.Name) + strings.Join(testChatTwo.Users, `", "`) + `"]}`

		apitest.New("InternalError").
			Handler(http.HandlerFunc(ch.Add)).
			Method("Post").
			URL(utils.GetAPIAddress("addChat")).
			Body(jsonBody).
			Expect(t).
			Status(http.StatusInternalServerError).
			End()

	})

	t.Run("MalformedJSON", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		apitest.New("MalformedJSON").
			Handler(http.HandlerFunc(ch.Add)).
			Method("Post").
			URL(utils.GetAPIAddress("addChat")).
			Body(fmt.Sprintf(`{"username": "%s"`, testUserOne.UserName)).
			Expect(t).
			Status(http.StatusInternalServerError).
			Assert(jsonpath.Contains(`$.message`, "Error unmarshaling json")).
			End()

	})
}


func TestGetUserChatsSorted(t *testing.T) {
	t.Run("ChatGetOK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := useCases.NewMockChatsUCInterface(ctrl)
		mockUseCase.EXPECT().GetUserChatsSorted(&testUserIdOne).Return(false, testUserOneChats, nil)

		ch.ChatsUC = mockUseCase

		jsonBody := fmt.Sprintf(`{"user": "%s"}`, testUserIdOne.UserId)

		apitest.New("ChatGetOK").
			Handler(http.HandlerFunc(ch.Get)).
			Method("Post").
			URL(utils.GetAPIAddress("getChats")).
			Body(jsonBody).
			Expect(t).
			Status(http.StatusOK).
			End()
	})

	t.Run("ChatGetUserNotExist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := useCases.NewMockChatsUCInterface(ctrl)
		mockUseCase.EXPECT().GetUserChatsSorted(&testUserIdWrong).Return(true, []models.Chat{}, errors.New("user doesn't exist"))

		ch.ChatsUC = mockUseCase

		jsonBody := fmt.Sprintf(`{"user": "%s"}`, testUserIdWrong.UserId)

		apitest.New("ChatGetUserNotExist").
			Handler(http.HandlerFunc(ch.Get)).
			Method("Post").
			URL(utils.GetAPIAddress("getChats")).
			Body(jsonBody).
			Expect(t).
			Status(http.StatusBadRequest).
			End()
	})

	t.Run("InternalError", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := useCases.NewMockChatsUCInterface(ctrl)
		mockUseCase.EXPECT().GetUserChatsSorted(&testUserIdOne).Return(false, []models.Chat{}, errors.New("database error"))

		ch.ChatsUC = mockUseCase

		jsonBody := fmt.Sprintf(`{"user": "%s"}`, testUserIdOne.UserId)

		apitest.New("InternalError").
			Handler(http.HandlerFunc(ch.Get)).
			Method("Post").
			URL(utils.GetAPIAddress("getChats")).
			Body(jsonBody).
			Expect(t).
			Status(http.StatusInternalServerError).
			End()

	})

	t.Run("MalformedJSON", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		apitest.New("MalformedJSON").
			Handler(http.HandlerFunc(ch.Get)).
			Method("Post").
			URL(utils.GetAPIAddress("getChats")).
			Body(fmt.Sprintf(`{"user" "%s"}`, testUserIdOne.UserId)).
			Expect(t).
			Status(http.StatusInternalServerError).
			Assert(jsonpath.Contains(`$.message`, "Error unmarshaling json")).
			End()

	})

}