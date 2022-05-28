package action

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/internal/repository"
	"github.com/mark-marushak/bot-english-book/logger"
	"math/rand"
	"reflect"
)

type AdaptorTelegramAction struct {
	BaseAction
}

func (a AdaptorTelegramAction) GetUpdate() *tgbotapi.Update {
	if update, ok := reflect.ValueOf(a.Data).Interface().(tgbotapi.Update); ok {
		return &update
	}

	return nil
}

func (a AdaptorTelegramAction) GetBotAPI() *tgbotapi.BotAPI {
	if bot, ok := reflect.ValueOf(a.Bot).Interface().(tgbotapi.BotAPI); ok {
		return &bot
	}

	return nil
}

func (AdaptorTelegramAction) updateStatusUser(chatID int64, bookID uint) error {
	repo := model.NewUserService(repository.NewUserRepository())
	user, err := repo.Get(model.User{ChatID: chatID})
	if err != nil {
		logger.Get().Error("Error while getting user for updating status")
		return err
	}

	user.Status = model.USER_STUDY
	user.BookID = bookID
	err = repo.Update(user)

	if err != nil {
		logger.Get().Error("Error while updating user status and book id")
		return err
	}

	return nil
}

func (a AdaptorTelegramAction) Lesson(book model.Book) (string, error) {
	random := func() int { return rand.Intn(len(book.Words)) }
	indexed := make(map[int]int)
	options := make([]string, 4)
	var index, order, correct int
	for len(indexed) < 4 {
		index = random()
		if _, ok := indexed[index]; ok {
			continue
		}
		if order == 0 {
			correct = index
		}
		indexed[index] = order
		order++
	}

	wordRepo := model.NewWordService(repository.NewWordRepository())
	for index, order := range indexed {
		word, err := wordRepo.GetTranslate(book.Words[index])
		if err != nil {
			logger.Get().Error("StudyStart: sorting words: %v", err)
			return "Сталась помилка, спробуй ще раз через 1 хвилину", nil
		}

		options[order] = word.Text
	}

	repo := model.NewUserService(repository.NewUserRepository())
	user, err := repo.Get(model.User{ChatID: a.GetUpdate().FromChat().ID})
	if err != nil {
		logger.Get().Error("Error while getting user for updating status")
		return "", err
	}

	a.GetBotAPI().Send(tgbotapi.NewStopPoll(a.GetUpdate().FromChat().ID, user.PollID))

	poll := tgbotapi.NewPoll(a.GetUpdate().FromChat().ID, fmt.Sprintf("Translate this %s", book.Words[correct].Text), options...)
	poll.CorrectOptionID = 0
	poll.Type = "quiz"
	poll.AllowsMultipleAnswers = true
	poll.IsAnonymous = false
	poll.OpenPeriod = 60

	msg, _ := a.GetBotAPI().Send(poll)
	user.PollID = msg.MessageID
	repo.Update(user)

	return "", nil
}
