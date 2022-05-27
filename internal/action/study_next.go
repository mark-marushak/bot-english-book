package action

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/internal/repository"
	"github.com/mark-marushak/bot-english-book/logger"
	"math/rand"
)

type StudyNext struct {
	AdaptorTelegramAction
}

func (s StudyNext) Keyboard(i ...interface{}) interface{} {
	return tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButtonContact("Головне меню")))
}

func (s StudyNext) Output(i ...interface{}) (string, error) {
	update := s.GetUpdate()
	userRepo := model.NewUserService(repository.NewUserRepository())
	user, err := userRepo.Get(model.User{ChatID: update.FromChat().ID})
	if err != nil {
		logger.Get().Error("StudyNext: getting user from db: %v", err)
	}

	bookRepo := model.NewBookService(repository.NewBookRepository())
	book, err := bookRepo.Get(model.Book{ID: user.BookID})
	if err != nil {
		logger.Get().Error("StudyNext: getting book by id: %v", err)
	}

	db.DB().Where(book).Preload("Words").Find(&book)

	if len(book.Words) <= 0 {
		return "Слова ще обробляються. Почекайте повідомлення про закінчення!", nil
	}

	random := func() int { return rand.Intn(len(book.Words)) }
	indexed := make(map[int]int)
	options := make([]string, 4)
	var index int
	for i := 0; i < 4; i++ {
		index = random()
		if _, ok := indexed[index]; ok {
			continue
		}
		indexed[index] = i
	}

	for index, order := range indexed {
		options[order] = book.Words[index].Text
	}

	poll := tgbotapi.NewPoll(s.GetChat(), fmt.Sprintf("Translate this %s", options[0]), options...)
	poll.CorrectOptionID = 0
	poll.Type = "quiz"
	poll.AllowsMultipleAnswers = true

	return "Ок, ось тобі завдання, в тебе є 1 хвилина", nil
}
