package action

import (
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/internal/repository/gorm"
	"github.com/mark-marushak/bot-english-book/logger"
)

type StudyStart struct {
	AdaptorTelegramAction
}

func (s StudyStart) Keyboard(i ...interface{}) interface{} {
	return NextLessonButton
}

func (s StudyStart) Output(i ...interface{}) (string, error) {
	update := s.GetUpdate()
	userRepo := model.NewUserService(gorm.NewUserRepository())
	user, err := userRepo.Get(model.User{ChatID: update.FromChat().ID})
	if err != nil {
		logger.Get().Error("StudyStart: getting user from db: %v", err)
	}

	bookRepo := model.NewBookService(gorm.NewBookRepository())
	book, err := bookRepo.Get(model.Book{ID: user.BookID})
	if err != nil {
		logger.Get().Error("StudyStart: getting book by id: %v", err)
	}

	db.Gorm().Where(book).Preload("Words").Find(&book)

	if len(book.Words) <= 0 {
		return "Слова ще обробляються. Почекайте повідомлення про закінчення!", nil
	}

	out, err := s.Lesson(book)
	if err != nil {
		return out, nil
	}

	return "Ок, ось тобі завдання, в тебе є 1 хвилина", nil
}
