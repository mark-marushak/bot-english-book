package action

import (
	"fmt"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/internal/repository"
	"strconv"
)

type BookLibrary struct {
	AdaptorTelegramAction
}

func (BookLibrary) GetInlineButton(id uint) interface{} {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Обрати", "book-id:"+strconv.Itoa(int(id))),
		),
	)
}

func (b BookLibrary) Keyboard(i ...interface{}) interface{} {
	return DoNothingButton
}

func (b BookLibrary) Output(i ...interface{}) (string, error) {
	bookService := model.NewBookService(repository.NewBookRepository())
	books, err := bookService.FindAll()
	if err != nil {
		return "", err
	}

	var book model.Book
	for i := 0; i < len(books); i++ {
		book = books[i]
		msg := tg.NewMessage(b.GetUpdate().FromChat().ID, fmt.Sprintf("%s\n", book.Name))
		msg.ReplyMarkup = b.GetInlineButton(book.ID)
		b.GetBotAPI().Send(msg)
	}

	return "Дивись та обирай", nil
}

//func (c ChooseBookHandler) Send(bot *tg.BotAPI, update tg.Update) (tg.Message, error) {
//	bookService := model.NewBookService(repository.NewBookRepository())
//	books, err := bookService.FindAll()
//	if err != nil {
//		return tg.Message{}, err
//	}
//
//	for i := 0; i < len(books); i++ {
//		var id string
//
//		id = base64.StdEncoding.EncodeToString([]byte(books[i].Name))
//		msg := tg.NewMessage(update.Message.Chat.ID, books[i].Name)
//		msg.ReplyMarkup = buttonChoseBookFunc(id)
//
//		//telegram.GetRouteMap().AddHandlerCallback(telegram.NewReplayRoute(id, telegram.NewHandler(ChooseBookReplay{})))
//
//		bot.Send(msg)
//	}
//
//	return tg.Message{}, nil
//}
//
//func (ChooseBookHandler) output(list []model.Book) string {
//	var b strings.Builder
//
//	for i := 0; i < len(list); i++ {
//		b.WriteString(fmt.Sprintf("%s\n", list[i].Name))
//	}
//
//	return b.String()
//}
