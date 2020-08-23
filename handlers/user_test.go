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
)

var uh UsersHandlers

var testUserOne = models.User{
	UserName: fake.UserName(),
}

var testUserTwo = models.User{
	UserName: fake.UserName(),
}

func TestAddUser(t *testing.T) {
	t.Run("UserAddOK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := useCases.NewMockUsersUCInterface(ctrl)
		mockUseCase.EXPECT().Add(&testUserOne).Return(false, nil)

		uh.UsersUC = mockUseCase

		apitest.New("UserAddOK").
			Handler(http.HandlerFunc(uh.Add)).
			Method("Post").
			URL(utils.GetAPIAddress("addUser")).
			Body(fmt.Sprintf(`{"username": "%s" }`, testUserOne.UserName)).
			Expect(t).
			Status(http.StatusCreated).
			Assert(jsonpath.Matches(`$.id`, `^[0-9]*$`)).
			End()
	})

	t.Run("UserAddExists", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := useCases.NewMockUsersUCInterface(ctrl)
		mockUseCase.EXPECT().Add(&testUserOne).Return(true, errors.New("already exists user error"))

		uh.UsersUC = mockUseCase

		apitest.New("UserAddExists").
			Handler(http.HandlerFunc(uh.Add)).
			Method("Post").
			URL(utils.GetAPIAddress("addUser")).
			Body(fmt.Sprintf(`{"username": "%s" }`, testUserOne.UserName)).
			Expect(t).
			Status(http.StatusConflict).
			End()

	})

	t.Run("InternalError", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := useCases.NewMockUsersUCInterface(ctrl)
		mockUseCase.EXPECT().Add(&testUserTwo).Return(false, errors.New("database error"))

		uh.UsersUC = mockUseCase

		apitest.New("InternalError").
			Handler(http.HandlerFunc(uh.Add)).
			Method("Post").
			URL(utils.GetAPIAddress("addUser")).
			Body(fmt.Sprintf(`{"username": "%s" }`, testUserTwo.UserName)).
			Expect(t).
			Status(http.StatusInternalServerError).
			End()

	})

	t.Run("MalformedJSON", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		apitest.New("MalformedJSON").
			Handler(http.HandlerFunc(uh.Add)).
			Method("Post").
			URL(utils.GetAPIAddress("addUser")).
			Body(fmt.Sprintf(`{"username": "%s"`, testUserOne.UserName)).
			Expect(t).
			Status(http.StatusInternalServerError).
			Assert(jsonpath.Contains(`$.message`, "Error unmarshaling json")).
			End()

	})
}
