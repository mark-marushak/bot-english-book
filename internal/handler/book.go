package handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type BookHanlder struct {
}

func (b BookHanlder) Send(bot *tgbotapi.BotAPI, update tgbotapi.Update) (tgbotapi.Message, error) {
	return tgbotapi.Message{}, nil
}
