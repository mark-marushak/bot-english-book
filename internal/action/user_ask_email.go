package action

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/internal/repository"
	"github.com/mark-marushak/bot-english-book/logger"
	"reflect"
)

type UserAskEmail struct {
	BaseAction
	ChatID int64
	Data   interface{}
}

const userAskEmailText = `Welcome on a board %s, 
			you can either upload one book or choose any book in the bot 
			to start study or just preparation to read full book`

func (u UserAskEmail) Keyboard(i ...interface{}) interface{} {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Upload Book"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Choose Book"),
		),
	)
}

func (u UserAskEmail) Output(i ...interface{}) string {
	data := u.GetData()
	var update tgbotapi.Update
	update, _ = reflect.ValueOf(data).Interface().(tgbotapi.Update)

	repo := model.NewUserService(repository.NewUserRepository())
	err := repo.Update(model.User{ChatID: update.FromChat().ID, Email: update.Message.Text})

	if err != nil {
		logger.Get().Error("error while saving a new user %v", err)
	}

	user, err := repo.Get(model.User{
		ChatID: update.SentFrom().ID,
	})

	return fmt.Sprintf(userSaveText, user.FirstName)
}
