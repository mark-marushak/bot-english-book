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

	education, err := userRepo.GetEducationByUserID(user.ID)
	if err != nil {
		return "", err
	}

	var count int
	err = db.Sqlx().Get(&count, `select count(*) from book_words where book_id = $1`, education.BookID)
	if err != nil {
		return "", err
	}

	if count <= 0 {
		return "Слова ще обробляються. Почекайте повідомлення про закінчення!", nil
	}

	out, err := s.Lesson(model.Book{ID: education.BookID})
	if err != nil {
		return out, nil
	}

	return "Ок, ось тобі завдання, в тебе є 1 хвилина", nil
}
