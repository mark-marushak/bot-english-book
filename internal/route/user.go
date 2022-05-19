package route

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-marushak/bot-english-book/internal/action"
	"github.com/mark-marushak/bot-english-book/pkg/telegram"
)

type UserRoute struct {
	Bot    *tgbotapi.BotAPI
	Update tgbotapi.Update
	action *telegram.ActionService
}

func (u *UserRoute) SetBot(bot *tgbotapi.BotAPI) {
	u.Bot = bot
}

func (u *UserRoute) SetUpdate(update tgbotapi.Update) {
	u.Update = update
}

var commands = map[string]telegram.ActionRepository{
	"start": &action.UserSave{},
}

var messages = map[string]telegram.ActionRepository{
	"phone": &action.UserPhone{},
}

var callbacks = map[string]telegram.ActionRepository{
	"start": &action.UserStudy{},
}

var contactRoutes = map[string]telegram.ActionRepository{
	"": &action.UserPhone{},
}

func (u UserRoute) find(list map[string]telegram.ActionRepository, text string) *telegram.ActionService {
	if found, ok := list[text]; ok {
		return telegram.NewAction(
			found,
		)
	}

	return nil
}

func (u *UserRoute) Analyze() (int64, error) {
	if u.Update.CallbackQuery != nil {
		u.action = u.find(callbacks, u.Update.CallbackData())
		return u.Update.CallbackQuery.Message.Chat.ID, nil
	}

	if u.Update.Message.IsCommand() {
		u.action = u.find(commands, u.Update.Message.Command())
		return u.Update.Message.Chat.ID, nil
	}

	if u.Update.Message.Contact != nil {
		u.action = u.find(contactRoutes, u.Update.Message.Text)
		return u.Update.Message.Chat.ID, nil
	}

	if u.Update.Message != nil {
		u.action = u.find(messages, u.Update.Message.Text)
		return u.Update.Message.Chat.ID, nil
	}

	return 0, telegram.NotFoundError
}

func (u *UserRoute) Response(chatID int64) (err error) {
	if u.action != nil {
		message := tgbotapi.NewMessage(chatID, u.action.Output())
		switch t := u.action.Keyboard().(type) {
		case tgbotapi.ReplyKeyboardMarkup:
			message.ReplyMarkup = t
		case tgbotapi.InlineKeyboardMarkup:
			message.ReplyMarkup = t
		}
		_, err = u.Bot.Send(message)
	}

	return
}
