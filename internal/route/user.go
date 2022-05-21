package route

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-marushak/bot-english-book/internal/action"
	"github.com/mark-marushak/bot-english-book/pkg/telegram"
	"regexp"
)

type UserRoute struct {
	Bot    *tgbotapi.BotAPI
	Update tgbotapi.Update
	action telegram.ActionService

	route map[string]map[string]telegram.ActionService
}

func (u *UserRoute) SetupRoutes() telegram.RouteService {
	u.route = map[string]map[string]telegram.ActionService{
		"messages": map[string]telegram.ActionService{
			emailPattern: &action.UserAskEmail{},
		},
		"commands": map[string]telegram.ActionService{
			"start": &action.StartHandler{},
		},
		"callbacks": map[string]telegram.ActionService{
			//"start": &action.UserStudy{},
		},
		"contact": map[string]telegram.ActionService{
			"": &action.UserAskPhone{},
		},
	}
	return u
}

var emailPattern = "^([a-zA-Z0-9_\\-\\.]+)@([a-zA-Z0-9_\\-\\.]+)\\.([a-zA-Z]{2,5})$"

func (u *UserRoute) SetBot(bot *tgbotapi.BotAPI) {
	u.Bot = bot
}

func (u *UserRoute) SetUpdate(update tgbotapi.Update) {
	u.Update = update
}

func (u UserRoute) find(list, text string) telegram.ActionService {

	for cond, found := range u.route["messages"] {
		if ok, _ := regexp.Match(cond, []byte(text)); ok {
			return telegram.NewAction(
				found,
			)
		}
	}

	if found, ok := u.route[list][text]; ok {
		return telegram.NewAction(
			found,
		)
	}

	return nil
}

func (u *UserRoute) Analyze() (int64, error) {
	if u.Update.CallbackQuery != nil {
		u.action = u.find("callbacks", u.Update.CallbackData())
		return u.Update.CallbackQuery.Message.Chat.ID, nil
	}

	if u.Update.Message.IsCommand() {
		u.action = u.find("commands", u.Update.Message.Command())
		return u.Update.Message.Chat.ID, nil
	}

	if u.Update.Message.Contact != nil {
		u.action = u.find("contact", u.Update.Message.Text)
		return u.Update.Message.Chat.ID, nil
	}

	if u.Update.Message != nil {
		u.action = u.find("messages", u.Update.Message.Text)
		return u.Update.Message.Chat.ID, nil
	}

	return 0, telegram.NotFoundError
}

func (u *UserRoute) Response(chatID int64) (err error) {
	if u.action != nil {
		u.action.SetData(u.Update)
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
