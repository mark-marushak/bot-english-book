package action

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/internal/repository"
	"github.com/mark-marushak/bot-english-book/logger"
	"reflect"
)

type UserAskPhone struct {
	BaseAction
}

const userAskPhoneText = `If change your phone number you will lose all your data. 
To recover your data We send you code on email to identify your new phone`

func (u UserAskPhone) Keyboard(i ...interface{}) interface{} {
	return tgbotapi.NewRemoveKeyboard(false)
}

func (u UserAskPhone) Output(i ...interface{}) string {
	data := u.GetData()
	var update tgbotapi.Update
	update, _ = reflect.ValueOf(data).Interface().(tgbotapi.Update)

	contact := update.Message.Contact

	repo := model.NewUserService(repository.NewUserRepository())
	err := repo.Create(model.User{
		ChatID:    update.SentFrom().ID,
		Phone:     contact.PhoneNumber,
		Email:     "",
		FirstName: contact.FirstName,
		Books:     nil,
		BookID:    0,
		Status:    model.USER_NEW,
	})

	if err != nil {
		logger.Get().Error("error while saving a new user %v", err)
	}

	return userAskPhoneText
}
