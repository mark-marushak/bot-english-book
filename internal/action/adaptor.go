package action

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/internal/repository/gorm"
	"github.com/mark-marushak/bot-english-book/internal/repository/sqlx"
	"github.com/mark-marushak/bot-english-book/logger"
	"math/rand"
	"reflect"
	"regexp"
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
	userService := model.NewUserService(gorm.NewUserRepository())
	user, err := userService.Get(model.User{ChatID: chatID})
	if err != nil {
		return err
	}

	bookService := model.NewBookService(gorm.NewBookRepository())
	book, err := bookService.Get(model.Book{ID: bookID})
	if err != nil {
		return err
	}

	educationRepository, err := sqlx.NewEducationRepository(user, book)
	if err != nil {
		return err
	}

	educationService := model.NewEducationService(educationRepository)

	get, err := educationService.Get(user.ID)
	if err != nil {
		return err
	}

	if get.UserID == 0 {
		err = educationService.CreateRelation()
		if err != nil {
			return err
		}
	} else {
		err = educationService.Update(user.ID, book.ID)
		if err != nil {
			return err
		}
	}

	user.Status = model.USER_STUDY
	_, err = userService.Update(user)
	if err != nil {
		logger.Get().Error("Error while updating user status and book id")
		return err
	}

	return nil
}

func (a AdaptorTelegramAction) Lesson(book model.Book) (string, error) {
	wordService := model.NewWordService(gorm.NewWordRepository())
	userService := model.NewUserService(gorm.NewUserRepository())
	user, err := userService.Get(model.User{ChatID: a.GetUpdate().FromChat().ID})
	if err != nil {
		logger.Get().Error("Error while getting user for updating status")
		return "", err
	}

	repository, err := sqlx.NewEducationRepository(user, book)
	if err != nil {
		return "", err
	}
	educationService := model.NewEducationService(repository)
	words, err := educationService.GetUnknownWords()
	if err != nil {
		return "", err
	}

	random := func() int { return rand.Intn(len(words)) }

	indexed := make(map[int]int)
	reverseIndexed := make(map[int]int)
	var index, order int
	for len(indexed) < 4 {
		index = random()
		if _, ok := indexed[index]; ok {
			continue
		}

		indexed[index] = order
		reverseIndexed[order] = index
		order++
	}

	possible := make(map[int]map[string]string)
	options := make([]string, 4)
	for index, order = range indexed {
		possible[order] = make(map[string]string)
		get, err := wordService.Get(model.Word{ID: words[index]})
		if err != nil {
			return "", err
		}

		translated, err := wordService.GetTranslate(get)
		if err != nil {
			logger.Get().Error("StudyStart: sorting words: %v", err)
			return "Сталась помилка, спробуй ще раз через 1 хвилину", nil
		}
		possible[order]["en"] = get.Text
		possible[order]["ua"] = translated.Text

		options[order] = translated.Text
	}

	getPoll, err := educationService.GetPoll()
	if err != nil {
		return "", err
	}

	if getPoll > 0 {
		_, err = a.GetBotAPI().Send(tgbotapi.NewStopPoll(a.GetUpdate().FromChat().ID, getPoll))
		if err != nil {
			if regexp.MustCompile("close").Match([]byte(err.Error())) == false {
				logger.Get().Error("Error while stopping old poll: %v", err)
				return "", err
			}
		}
	}

	correct := rand.Intn(4)

	poll := tgbotapi.NewPoll(a.GetUpdate().FromChat().ID, fmt.Sprintf("Перекладіть слово: %s", possible[correct]["en"]), options...)
	poll.CorrectOptionID = int64(correct)
	poll.Type = "quiz"
	poll.AllowsMultipleAnswers = true
	poll.IsAnonymous = false
	poll.OpenPeriod = 60

	msg, _ := a.GetBotAPI().Send(poll)
	err = educationService.SetPoll(msg.MessageID, correct, words[reverseIndexed[correct]])
	if err != nil {
		return "", err
	}

	return "", nil
}
