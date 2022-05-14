package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ServiceHandler struct {
	repository RepositoryHandler
}

type RepositoryHandler interface {
	Send(bot *tgbotapi.BotAPI, update tgbotapi.Update) (tgbotapi.Message, error)
}

func NewHandler(repository RepositoryHandler) *ServiceHandler {
	return &ServiceHandler{
		repository: repository,
	}
}

func (h ServiceHandler) Send(bot *tgbotapi.BotAPI, update tgbotapi.Update) (tgbotapi.Message, error) {
	return h.repository.Send(bot, update)
}
