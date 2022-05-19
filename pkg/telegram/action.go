package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

/*
ActionService implement three basic functions
Send - resposible for collect keyboard and text for sending
Keyboard - return keyboard based on some data from message
Output - return message text prepared for sending
*/
type ActionRepository interface {
	Keyboard(i ...interface{}) interface{}
	Output(...interface{}) string
}

type Action struct {
	chatID int64
}

func (a *Action) SetChat(i int64) {
	a.chatID = i
}

func (a Action) GetChat() int64 {
	return a.chatID
}

type ActionService struct {
	Update tgbotapi.Update
	Bot    tgbotapi.BotAPI
	action ActionRepository
}

func NewAction(action ActionRepository) *ActionService {
	return &ActionService{
		action: action,
	}
}

func (a ActionService) Keyboard(i ...interface{}) interface{} {
	return a.action.Keyboard(i)
}

func (a ActionService) Output(i ...interface{}) string {
	return a.action.Output(i)
}
