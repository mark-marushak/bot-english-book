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
	action *telegram.ActionService
}

var massages = map[string]telegram.ActionRepository{
	"choose-book": &action.BookChoose{},
}

func (u BookRoute) find(list map[string]telegram.ActionRepository, text string) *telegram.ActionService {
	if found, ok := list[text]; ok {
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

		b.action = b.find(messages, text)
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
