package action

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-marushak/bot-english-book/pkg/telegram"
)

type UserSave struct {
	telegram.Action
}

const userSaveText = `If you want use this bot, 
you should share you phone number to active
Trial subscribe`

func (u UserSave) Keyboard(i ...interface{}) interface{} {
	return tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButtonContact("Phone Number")))
}

func (u UserSave) Output(i ...interface{}) string {
	return userSaveText
}
