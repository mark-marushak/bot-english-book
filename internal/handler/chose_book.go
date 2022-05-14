package handler

import (
	"encoding/base64"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-marushak/bot-english-book/internal/mock"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/pkg/telegram"
	"strings"
)

type ChooseBookHandler struct {
}

var buttonChoseBookFunc = func(id string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("choose", id),
		),
	)
}

func (c ChooseBookHandler) Send(bot *tgbotapi.BotAPI, update tgbotapi.Update) (tgbotapi.Message, error) {
	bookService := model.NewBookService(repository.NewBookRepository())
	books, err := bookService.FindAll()
	if err != nil {
		return tgbotapi.Message{}, err
	}

	for i := 0; i < len(books); i++ {
		var id string

		id = base64.StdEncoding.EncodeToString([]byte(books[i].Name))
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, books[i].Name)
		msg.ReplyMarkup = buttonChoseBookFunc(id)

		telegram.GetRouteMap().AddHandlerCallback(telegram.NewReplayRoute(id, telegram.NewHandler(ChooseBookReplay{})))

		bot.Send(msg)
	}

	return tgbotapi.Message{}, nil
}

func (ChooseBookHandler) output(list []model.Book) string {
	var b strings.Builder

	for i := 0; i < len(list); i++ {
		b.WriteString(fmt.Sprintf("%s\n", list[i].Name))
	}

	return b.String()
}

type ChooseBookReplay struct{}

func (c ChooseBookReplay) Send(bot *tgbotapi.BotAPI, update tgbotapi.Update) (tgbotapi.Message, error) {
	bookService := model.NewBookService(repository.NewBookRepository())

	name, err := base64.StdEncoding.DecodeString(update.CallbackQuery.Data)
	book, err := bookService.FindByName(string(name))
	if err != nil {
		return tgbotapi.Message{}, err
	}

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, book.Name)
	return bot.Send(msg)
}

func (ChooseBookReplay) output(list []model.Book) string {
	var b strings.Builder

	for i := 0; i < len(list); i++ {
		b.WriteString(fmt.Sprintf("%s\n", list[i].Name))
	}

	return b.String()
}
