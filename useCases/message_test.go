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

var testMessageOne = models.Message{
	Chat:   "3",
	Author: "2",
	Text:   fake.Sentence(),
}

var testMessageOneChatId = 3

var testMessageWrong = models.Message{
	Chat:   "5",
	Author: "2",
	Text:   fake.Sentence(),
}

var testChatIdOne = models.ChatId{
	ChatId: "1",
}

var testChatIdOneId = 1

var testChatIdWrong = models.ChatId{
	ChatId: "15",
}

var testChatOneMessages = []models.Message{
	{Id: "2", Chat: "1", Author: "1", Text: "sss", Created: time.Now()},
	{Id: "5", Chat: "1", Author: "2", Text: "sdfeeei", Created: time.Now()},
	{Id: "6", Chat: "1", Author: "2", Text: "sdfeeei", Created: time.Now()},
	{Id: "7", Chat: "1", Author: "2", Text: "sdfeeei", Created: time.Now()},
}

func TestAddMessage(t *testing.T) {
	t.Run("MessageAddOK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepoChats := repository.NewMockChatsRepoInterface(ctrl)
		mockRepoChats.EXPECT().CheckUserInChat(testMessageOne.Author, testMessageOne.Chat).Return(testMessageOneChatId, nil)

		mockRepoMessages := repository.NewMockMessagesRepoInterface(ctrl)
		mockRepoMessages.EXPECT().InsertMessage(&testMessageOne).Return(utils.NO_ERROR, nil)

		messagesUseCase := MessagesUC{
			ChatsRepo:    mockRepoChats,
			MessagesRepo: mockRepoMessages,
		}

		userError, err := messagesUseCase.Add(&testMessageOne)

		assert.NoError(t, err)
		assert.Equal(t, userError, false)
	})

	t.Run("MessageAddUserNotInChat", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepoChats := repository.NewMockChatsRepoInterface(ctrl)
		mockRepoChats.EXPECT().CheckUserInChat(testMessageOne.Author, testMessageOne.Chat).Return(utils.ERROR_ID, nil)

		mockRepoMessages := repository.NewMockMessagesRepoInterface(ctrl)

		messagesUseCase := MessagesUC{
			ChatsRepo:    mockRepoChats,
			MessagesRepo: mockRepoMessages,
		}

		userError, err := messagesUseCase.Add(&testMessageOne)

		assert.Error(t, err)
		assert.Equal(t, userError, true)
		assert.Equal(t, err, fmt.Errorf("this user is not in this chat"))
	})

	t.Run("DBErrorFirst", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepoChats := repository.NewMockChatsRepoInterface(ctrl)
		mockRepoChats.EXPECT().CheckUserInChat(testMessageOne.Author, testMessageOne.Chat).Return(testMessageOneChatId, errors.New("database error"))

		mockRepoMessages := repository.NewMockMessagesRepoInterface(ctrl)

		messagesUseCase := MessagesUC{
			ChatsRepo:    mockRepoChats,
			MessagesRepo: mockRepoMessages,
		}

		userError, err := messagesUseCase.Add(&testMessageOne)

		assert.Error(t, err)
		assert.Equal(t, userError, false)
	})

	t.Run("DBErrorSecond", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepoChats := repository.NewMockChatsRepoInterface(ctrl)
		mockRepoChats.EXPECT().CheckUserInChat(testMessageOne.Author, testMessageOne.Chat).Return(testMessageOneChatId, nil)

		mockRepoMessages := repository.NewMockMessagesRepoInterface(ctrl)
		mockRepoMessages.EXPECT().InsertMessage(&testMessageOne).Return(utils.SERVER_ERROR, errors.New("database error"))
		messagesUseCase := MessagesUC{
			ChatsRepo:    mockRepoChats,
			MessagesRepo: mockRepoMessages,
		}

		userError, err := messagesUseCase.Add(&testMessageOne)

		assert.Error(t, err)
		assert.Equal(t, userError, false)
	})

	t.Run("UserError", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepoChats := repository.NewMockChatsRepoInterface(ctrl)
		mockRepoChats.EXPECT().CheckUserInChat(testMessageOne.Author, testMessageOne.Chat).Return(testMessageOneChatId, nil)

		mockRepoMessages := repository.NewMockMessagesRepoInterface(ctrl)
		mockRepoMessages.EXPECT().InsertMessage(&testMessageOne).Return(utils.USER_ERROR, errors.New("user error"))
		messagesUseCase := MessagesUC{
			ChatsRepo:    mockRepoChats,
			MessagesRepo: mockRepoMessages,
		}

		userError, err := messagesUseCase.Add(&testMessageOne)

		assert.Error(t, err)
		assert.Equal(t, userError, true)
	})
}

func TestGetChatMessagesSorted(t *testing.T) {
	t.Run("MessagesGetOK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepoChats := repository.NewMockChatsRepoInterface(ctrl)
		mockRepoChats.EXPECT().CheckChat(testChatIdOne.ChatId).Return(testChatIdOneId, nil)

		mockRepoMessages := repository.NewMockMessagesRepoInterface(ctrl)
		mockRepoMessages.EXPECT().GetMessagesByChatId(testChatIdOne.ChatId).Return(testChatOneMessages, nil)

		messagesUseCase := MessagesUC{
			ChatsRepo:    mockRepoChats,
			MessagesRepo: mockRepoMessages,
		}

		chatNotExist, messages, err := messagesUseCase.GetChatMessagesSorted(&testChatIdOne)

		assert.NoError(t, err)
		assert.Equal(t, chatNotExist, false)
		assert.Equal(t, messages, testChatOneMessages)
	})

	t.Run("ChatNotExist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepoChats := repository.NewMockChatsRepoInterface(ctrl)
		mockRepoChats.EXPECT().CheckChat(testChatIdWrong.ChatId).Return(utils.ERROR_ID, nil)

		mockRepoMessages := repository.NewMockMessagesRepoInterface(ctrl)

		messagesUseCase := MessagesUC{
			ChatsRepo:    mockRepoChats,
			MessagesRepo: mockRepoMessages,
		}

		chatNotExist, messages, err := messagesUseCase.GetChatMessagesSorted(&testChatIdWrong)

		assert.Error(t, err)
		assert.Equal(t, chatNotExist, true)
		assert.Equal(t, messages, []models.Message{})
		assert.Equal(t, err, fmt.Errorf("this chat doesn't exist"))
	})

	t.Run("DBErrorFirst", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepoChats := repository.NewMockChatsRepoInterface(ctrl)
		mockRepoChats.EXPECT().CheckChat(testChatIdOne.ChatId).Return(testChatIdOneId, errors.New("database error"))

		mockRepoMessages := repository.NewMockMessagesRepoInterface(ctrl)

		messagesUseCase := MessagesUC{
			ChatsRepo:    mockRepoChats,
			MessagesRepo: mockRepoMessages,
		}

		chatNotExist, messages, err := messagesUseCase.GetChatMessagesSorted(&testChatIdOne)

		assert.Error(t, err)
		assert.Equal(t, chatNotExist, false)
		assert.Equal(t, messages, []models.Message{})
	})

	t.Run("DBErrorSecond", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepoChats := repository.NewMockChatsRepoInterface(ctrl)
		mockRepoChats.EXPECT().CheckChat(testChatIdOne.ChatId).Return(testChatIdOneId, nil)

		mockRepoMessages := repository.NewMockMessagesRepoInterface(ctrl)
		mockRepoMessages.EXPECT().GetMessagesByChatId(testChatIdOne.ChatId).Return([]models.Message{}, errors.New("database error"))

		messagesUseCase := MessagesUC{
			ChatsRepo:    mockRepoChats,
			MessagesRepo: mockRepoMessages,
		}

		chatNotExist, messages, err := messagesUseCase.GetChatMessagesSorted(&testChatIdOne)

		assert.Error(t, err)
		assert.Equal(t, chatNotExist, false)
		assert.Equal(t, messages, []models.Message{})
	})
}
