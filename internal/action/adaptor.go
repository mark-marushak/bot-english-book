package action

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"reflect"
)

type AdaptorTelegramAction struct {
	BaseAction
}

func (a AdaptorTelegramAction) GetUpdate() *tgbotapi.Update {
	if update, ok := reflect.ValueOf(a.Data).Interface().(tgbotapi.Update); ok {
		return &update
	}

	return nil
}

func (a AdaptorTelegramAction) GetBotAPI() *tgbotapi.BotAPI {
	if bot, ok := reflect.ValueOf(a.Bot).Interface().(tgbotapi.BotAPI); ok {
		return &bot
	}

	return nil
}
