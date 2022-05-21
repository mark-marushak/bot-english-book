package action

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StartHandler struct {
	BaseAction
}

const (
	startText = `Hello this bot help you to start read books without any problems!
				 Let's start with choosing a book you like!
				 Perhaps you dream about book we don't have, then click upload ahead!

				If you want use this bot, 
				you should share you phone number to active
				Trial subscribe
				`
)

func (s StartHandler) Keyboard(i ...interface{}) interface{} {
	return tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButtonContact("Phone Number")))
}

func (s StartHandler) Output(i ...interface{}) string {
	return startText
}
