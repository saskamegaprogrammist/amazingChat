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
	"time"
)

var testChatOne = models.Chat{
	Name:  fake.ProductName(),
	Users: []string{"1", "2"},
}

var testChatWrong = models.Chat{
	Name:  fake.ProductName(),
	Users: []string{"1", "2", "15"},
}
var testChatOneId = 1

var testUserIdOne = models.UserId{
	UserId: "1",
}

var testUserIdOneId = 1

var testUserIdWrong = models.UserId{
	UserId: "15",
}

var testUserOneChats = []models.Chat{
	{Id: "1", Name: "chat_1", Users: []string{"1", "2", "3"}, Created: time.Now()},
	{Id: "4", Name: "chat_3", Users: []string{"1", "2"}, Created: time.Now()},
	{Id: "3", Name: "chat_2", Users: []string{"1", "2", "3"}, Created: time.Now()},
}

func TestAddChat(t *testing.T) {
	t.Run("ChatAddOK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepoUsers := repository.NewMockUsersRepoInterface(ctrl)

		mockRepoChats := repository.NewMockChatsRepoInterface(ctrl)
		mockRepoChats.EXPECT().GetChatIdByName(&testChatOne).Return(utils.ERROR_ID, nil)
		mockRepoChats.EXPECT().InsertChat(&testChatOne).Return(utils.NO_ERROR, nil)

		chatsUseCase := ChatsUC{
			UsersRepo: mockRepoUsers,
			ChatsRepo: mockRepoChats,
		}

		chatExists, userError, err := chatsUseCase.Add(&testChatOne)

		assert.NoError(t, err)
		assert.Equal(t, chatExists, false)
		assert.Equal(t, userError, false)
	})

	t.Run("ChatAddExists", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepoUsers := repository.NewMockUsersRepoInterface(ctrl)

		mockRepoChats := repository.NewMockChatsRepoInterface(ctrl)
		mockRepoChats.EXPECT().GetChatIdByName(&testChatOne).Return(testChatOneId, nil)

		chatsUseCase := ChatsUC{
			UsersRepo: mockRepoUsers,
			ChatsRepo: mockRepoChats,
		}

		chatExists, userError, err := chatsUseCase.Add(&testChatOne)

		assert.Error(t, err)
		assert.Equal(t, chatExists, true)
		assert.Equal(t, userError, false)
		assert.Equal(t, err, fmt.Errorf("this name is already taken"))

	})

	t.Run("DBErrorFirst", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepoUsers := repository.NewMockUsersRepoInterface(ctrl)

		mockRepoChats := repository.NewMockChatsRepoInterface(ctrl)
		mockRepoChats.EXPECT().GetChatIdByName(&testChatOne).Return(utils.ERROR_ID, errors.New("database error"))

		chatsUseCase := ChatsUC{
			UsersRepo: mockRepoUsers,
			ChatsRepo: mockRepoChats,
		}

		chatExists, userError, err := chatsUseCase.Add(&testChatOne)

		assert.Error(t, err)
		assert.Equal(t, chatExists, false)
		assert.Equal(t, userError, false)
	})

	t.Run("DBErrorSecond", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepoUsers := repository.NewMockUsersRepoInterface(ctrl)

		mockRepoChats := repository.NewMockChatsRepoInterface(ctrl)
		mockRepoChats.EXPECT().GetChatIdByName(&testChatOne).Return(utils.ERROR_ID, nil)
		mockRepoChats.EXPECT().InsertChat(&testChatOne).Return(utils.SERVER_ERROR, errors.New("database error"))

		chatsUseCase := ChatsUC{
			UsersRepo: mockRepoUsers,
			ChatsRepo: mockRepoChats,
		}

		chatExists, userError, err := chatsUseCase.Add(&testChatOne)

		assert.Error(t, err)
		assert.Equal(t, chatExists, false)
		assert.Equal(t, userError, false)
	})

	t.Run("UserError", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepoUsers := repository.NewMockUsersRepoInterface(ctrl)

		mockRepoChats := repository.NewMockChatsRepoInterface(ctrl)
		mockRepoChats.EXPECT().GetChatIdByName(&testChatWrong).Return(utils.ERROR_ID, nil)
		mockRepoChats.EXPECT().InsertChat(&testChatWrong).Return(utils.USER_ERROR, errors.New("not existing users"))

		chatsUseCase := ChatsUC{
			UsersRepo: mockRepoUsers,
			ChatsRepo: mockRepoChats,
		}

		chatExists, userError, err := chatsUseCase.Add(&testChatWrong)

		assert.Error(t, err)
		assert.Equal(t, chatExists, false)
		assert.Equal(t, userError, true)
	})
}

func TestGetUserChatsSorted(t *testing.T) {
	t.Run("ChatsGetOK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepoUsers := repository.NewMockUsersRepoInterface(ctrl)
		mockRepoUsers.EXPECT().CheckUser(testUserIdOne.UserId).Return(testUserIdOneId, nil)

		mockRepoChats := repository.NewMockChatsRepoInterface(ctrl)
		mockRepoChats.EXPECT().GetChatsByUserId(testUserIdOne.UserId).Return(testUserOneChats, nil)

		chatsUseCase := ChatsUC{
			UsersRepo: mockRepoUsers,
			ChatsRepo: mockRepoChats,
		}

		userNotExist, chats, err := chatsUseCase.GetUserChatsSorted(&testUserIdOne)

		assert.NoError(t, err)
		assert.Equal(t, userNotExist, false)
		assert.Equal(t, chats, testUserOneChats)
	})

	t.Run("UserNotExist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepoUsers := repository.NewMockUsersRepoInterface(ctrl)
		mockRepoUsers.EXPECT().CheckUser(testUserIdWrong.UserId).Return(utils.ERROR_ID, nil)

		mockRepoChats := repository.NewMockChatsRepoInterface(ctrl)

		chatsUseCase := ChatsUC{
			UsersRepo: mockRepoUsers,
			ChatsRepo: mockRepoChats,
		}

		userNotExist, chats, err := chatsUseCase.GetUserChatsSorted(&testUserIdWrong)

		assert.Error(t, err)
		assert.Equal(t, userNotExist, true)
		assert.Equal(t, err, fmt.Errorf("this user doesn't exist"))
		assert.Equal(t, chats, []models.Chat{})
	})

	t.Run("DBErrorFirst", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepoUsers := repository.NewMockUsersRepoInterface(ctrl)
		mockRepoUsers.EXPECT().CheckUser(testUserIdOne.UserId).Return(testUserIdOneId, errors.New("database error"))

		mockRepoChats := repository.NewMockChatsRepoInterface(ctrl)

		chatsUseCase := ChatsUC{
			UsersRepo: mockRepoUsers,
			ChatsRepo: mockRepoChats,
		}

		userNotExist, chats, err := chatsUseCase.GetUserChatsSorted(&testUserIdOne)

		assert.Error(t, err)
		assert.Equal(t, userNotExist, false)
		assert.Equal(t, chats, []models.Chat{})
	})

	t.Run("DBErrorSecond", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepoUsers := repository.NewMockUsersRepoInterface(ctrl)
		mockRepoUsers.EXPECT().CheckUser(testUserIdOne.UserId).Return(testUserIdOneId, nil)

		mockRepoChats := repository.NewMockChatsRepoInterface(ctrl)
		mockRepoChats.EXPECT().GetChatsByUserId(testUserIdOne.UserId).Return([]models.Chat{}, errors.New("database error"))

		chatsUseCase := ChatsUC{
			UsersRepo: mockRepoUsers,
			ChatsRepo: mockRepoChats,
		}

		userNotExist, chats, err := chatsUseCase.GetUserChatsSorted(&testUserIdOne)

		assert.Error(t, err)
		assert.Equal(t, userNotExist, false)
		assert.Equal(t, chats, []models.Chat{})
	})
}
