package route

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-marushak/bot-english-book/internal/action"
	"github.com/mark-marushak/bot-english-book/pkg/telegram"
	"strings"
)

type BookRoute struct {
	Bot    *tgbotapi.BotAPI
	Update tgbotapi.Update
	action telegram.ActionService

	route map[string]map[string]telegram.ActionService
}

func (b *BookRoute) SetRoute() {
	b.route = map[string]map[string]telegram.ActionService{
		"messages": map[string]telegram.ActionService{
			"choose-book": &action.BookChoose{},
		},
	}
}

func (b BookRoute) find(list, text string) telegram.ActionService {
	if found, ok := b.route[list][text]; ok {
		return telegram.NewAction(
			found,
		)
	}

	return nil
}

func (b *BookRoute) Analyze() (int64, error) {
	if b.Update.Message != nil {

		text := b.Update.Message.Text
		text = strings.ToLower(text)
		text = strings.ReplaceAll(text, " ", "-")

		b.action = b.find("messages", text)
		return b.Update.FromChat().ID, nil
	}

	return 0, telegram.NotFoundError
}

func (b *BookRoute) Response(chatID int64) (err error) {
	if b.action != nil {
		message := tgbotapi.NewMessage(chatID, b.action.Output())
		switch t := b.action.Keyboard().(type) {
		case tgbotapi.ReplyKeyboardMarkup:
			message.ReplyMarkup = t
		}
		_, err = b.Bot.Send(message)
	}

	return
}
