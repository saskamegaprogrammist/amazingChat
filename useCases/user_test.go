package useCases

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/jackc/fake"
	"github.com/saskamegaprogrammist/amazingChat/models"
	"github.com/saskamegaprogrammist/amazingChat/repository"
	"github.com/saskamegaprogrammist/amazingChat/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testUserOne = models.User{
	Id:       "1",
	UserName: fake.UserName(),
}

var testUserTwo = models.User{
	Id:       "3",
	UserName: "three",
}

var testUserOneId = 1

func TestAddUser(t *testing.T) {
	t.Run("UserAddOK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := repository.NewMockUsersRepoInterface(ctrl)
		mockRepo.EXPECT().GetUserIdByUsername(&testUserOne).Return(utils.ERROR_ID, nil)
		mockRepo.EXPECT().InsertUser(&testUserOne).Return(nil)

		usersUseCase := UsersUC{
			UsersRepo: mockRepo,
		}
		exists, err := usersUseCase.Add(&testUserOne)

		assert.NoError(t, err)
		assert.Equal(t, exists, false)
	})

	t.Run("UserAddExists", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := repository.NewMockUsersRepoInterface(ctrl)
		mockRepo.EXPECT().GetUserIdByUsername(&testUserOne).Return(testUserOneId, nil)

		usersUseCase := UsersUC{
			UsersRepo: mockRepo,
		}
		exists, err := usersUseCase.Add(&testUserOne)

		assert.Error(t, err)
		assert.Equal(t, exists, true)
		assert.Equal(t, err, fmt.Errorf("this username is already taken"))
	})

	t.Run("DBError", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := repository.NewMockUsersRepoInterface(ctrl)
		mockRepo.EXPECT().GetUserIdByUsername(&testUserOne).Return(utils.ERROR_ID, errors.New("database error"))

		usersUseCase := UsersUC{
			UsersRepo: mockRepo,
		}
		exists, err := usersUseCase.Add(&testUserOne)

		assert.Error(t, err)
		assert.Equal(t, exists, false)

	})
}
