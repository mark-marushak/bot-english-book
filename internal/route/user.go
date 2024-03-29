package route

import (
	"github.com/mark-marushak/bot-english-book/internal/action"
	"github.com/mark-marushak/bot-english-book/pkg/telegram"
	"regexp"
)

var emailPattern = "^([a-zA-Z0-9_\\-\\.]+)@([a-zA-Z0-9_\\-\\.]+)\\.([a-zA-Z]{2,5})$"
var bookIDPattern = "^(book-id:)(\\d+)"

type UserRoute struct {
	baseRoute
}

func (u *UserRoute) SetupRoutes() telegram.RouteService {
	u.route = map[string]map[string]telegram.ActionService{
		"regex": {
			emailPattern:  &action.UserAskEmail{},
			bookIDPattern: &action.BookChoose{},
		},
		"messages": {
			action.StartStudy:             &action.StudyStart{},
			action.NextLesson:             &action.StudyStart{},
			action.BackToMainMenu:         &action.MainMenu{},
			action.DisplayStatisticButton: &action.DisplayStatistic{},
		},
		"commands": {
			"start": &action.StartHandler{},
		},
		"callbacks": {
			//"start": &action.UserStudy{},
		},
		"contact": {
			"": &action.UserAskPhone{},
		},
		"pollAnswer": {
			"": &action.PollAnswer{},
		},
	}
	return u
}

func (u UserRoute) RegexSearch(text string) (telegram.ActionService, error) {
	for cond, found := range u.route["regex"] {
		if ok, _ := regexp.Match(cond, []byte(text)); ok {
			return telegram.NewAction(
				found,
			), nil
		}
	}

	return nil, telegram.RouteNotFoundError
}

func (u UserRoute) MessageSearch(text string) (telegram.ActionService, error) {
	//text = strings.ToLower(text)
	//text = strings.ReplaceAll(text, " ", "-")

	if found, ok := u.route["messages"][text]; ok {
		return telegram.NewAction(
			found,
		), nil
	}

	return nil, telegram.RouteNotFoundError
}

func (u UserRoute) find(list, text string) (telegram.ActionService, error) {
	if found, ok := u.route[list][text]; ok {
		return telegram.NewAction(
			found,
		), nil
	}

	return nil, telegram.RouteNotFoundError
}

func (u *UserRoute) Analyze() (chatID int64, err error) {

	if u.Update.PollAnswer != nil {
		u.action, err = u.find("pollAnswer", "")
		chatID = u.Update.PollAnswer.User.ID
		return
	}

	if u.Update.FromChat() == nil {
		return 0, telegram.RouteNotFoundError
	}

	chatID = u.Update.FromChat().ID

	if u.Update.CallbackQuery != nil {
		u.action, err = u.RegexSearch(u.Update.CallbackData())
		if err != nil {
			u.action, err = u.find("callbacks", u.Update.CallbackData())
		}
		return
	}

	if u.Update.Message.IsCommand() {
		u.action, err = u.find("commands", u.Update.Message.Command())
		return
	}

	if u.Update.Message.Contact != nil {
		u.action, err = u.find("contact", u.Update.Message.Text)
		return
	}

	if u.Update.Message != nil {
		u.action, err = u.RegexSearch(u.Update.Message.Text)
		if err != nil {
			u.action, err = u.MessageSearch(u.Update.Message.Text)
		}

		return
	}

	return 0, telegram.RouteNotFoundError
}
