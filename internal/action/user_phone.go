package action

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-marushak/bot-english-book/pkg/telegram"
)

type UserPhone struct {
	telegram.Action
}

func (u UserPhone) Keyboard(i ...interface{}) interface{} {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Upload Book"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Choose Book"),
		),
	)

}

func (u UserPhone) Output(i ...interface{}) string {
	return "Thanks you and Welcome to a board!"
}
