package handler

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UserHandler struct{}

func (u UserHandler) Send(bot *tgbotapi.BotAPI, update tgbotapi.Update) (err error) {
	message := tgbotapi.NewMessage(update.Message.Chat.ID, "hello you are at home")

	_, err = bot.Send(message)
	if err != nil {
		err = fmt.Errorf("[Handler ERR]: %v ", err)
	}

	return err
}
