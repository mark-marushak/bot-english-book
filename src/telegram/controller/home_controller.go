package controller

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-marushak/bot-english-book/src/telegram"
)

type Home struct {
	next telegram.Handler
}

func (this Home) setNext(handler telegram.Handler) {
	this.next = handler
}

func (this Home) execute(update *tgbotapi.Update) {

	this.next.Execute(update)
}
