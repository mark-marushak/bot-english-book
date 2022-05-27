package route

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-marushak/bot-english-book/pkg/telegram"
)

type baseRoute struct {
	Bot    tgbotapi.BotAPI
	Update tgbotapi.Update
	action telegram.ActionService

	route map[string]map[string]telegram.ActionService
}

func (b baseRoute) Response(chatID int64) (err error) {
	if b.action != nil {
		b.action.SetData(b.Update)
		b.action.SetBot(b.Bot)

		output, err := b.action.Output()
		if err != nil {
			return err
		}

		message := tgbotapi.NewMessage(chatID, output)
		switch t := b.action.Keyboard().(type) {
		case tgbotapi.ReplyKeyboardMarkup:
			message.ReplyMarkup = t
		case tgbotapi.InlineKeyboardMarkup:
			message.ReplyMarkup = t
		}

		_, err = b.Bot.Send(message)
	}

	return
}

func (b *baseRoute) SetBot(bot tgbotapi.BotAPI) {
	b.Bot = bot
}

func (b *baseRoute) SetUpdate(update tgbotapi.Update) {
	b.Update = update
}
