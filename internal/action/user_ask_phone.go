package action

import (
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/internal/repository/gorm"
	"github.com/mark-marushak/bot-english-book/logger"
)

type UserAskPhone struct {
	AdaptorTelegramAction
}

//const userAskPhoneText = `If change your phone number you will lose all your data.
//To recover your data We send you code on email to identify your new phone`
const userAskPhoneText = `Як я і казав зараз треба відправити пошту, 
я не буду її зараз перевіряти тому будь певним що це саме твоя пошта!`

func (u UserAskPhone) Keyboard(i ...interface{}) interface{} {
	return DoNothingButton
}

func (u UserAskPhone) Output(i ...interface{}) (string, error) {

	contact := u.GetUpdate().Message.Contact

	repo := model.NewUserService(gorm.NewUserRepository())
	err := repo.Create(model.User{
		ChatID:    u.GetUpdate().SentFrom().ID,
		Phone:     contact.PhoneNumber,
		Email:     "",
		FirstName: contact.FirstName,
		Status:    model.USER_NEW,
	})

	if err != nil {
		logger.Get().Error("UserAskPhone: creating new user %v", err)
		return "", err
	}

	return userAskPhoneText, nil
}
