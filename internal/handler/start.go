package handler

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StartHandler struct {
}

const (
	startText = `Hello this bot help you to start read books without any problems!
				 Let's start with choosing a book you like!
				 Perhaps you dream about book we don't have, then click upload ahead!
				`
)

var keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Choose book"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Upload book"),
	),
)

func (s StartHandler) Send(bot *tgbotapi.BotAPI, update tgbotapi.Update) (err error) {
	message := tgbotapi.NewMessage(update.Message.Chat.ID, startText)

	message.ReplyMarkup = keyboard

	_, err = bot.Send(message)
	if err != nil {
		err = fmt.Errorf("[Handler ERR]: %v ", err)
	}

	return err
}
