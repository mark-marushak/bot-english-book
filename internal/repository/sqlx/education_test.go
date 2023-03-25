package sqlx

import (
	"github.com/mark-marushak/bot-english-book/config"
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/internal/repository/gorm"
	"github.com/mark-marushak/bot-english-book/logger"
	"testing"
)

func TestEducationRepository(t *testing.T) {
	config.NewConfig()
	logger.StartLogger()
	db.PrepareTable()

	var err error
	testUser := model.User{
		ChatID:    123,
		Phone:     "123123123",
		Email:     "mm@gmail.com",
		FirstName: "test user",
		Status:    "new",
	}

	testBook := model.Book{
		Name:   "testBOok",
		Path:   "/adfas/ss",
		Status: model.BOOK_UPLOAD,
	}

	failFunc := func(t *testing.T) {
		db.Sqlx().Queryx("delete from books where name = $1", testBook.Name)
		db.Sqlx().Queryx("delete from users where chat_id = $1", testUser.ChatID)
		db.Sqlx().Queryx("delete from educations where user_id = $1 and book_id = $2", testUser.ID, testBook.ID)
		t.FailNow()
	}

	clearFunc := func() {
		_, err = db.Sqlx().Queryx("delete from books where name = $1;", testBook.Name)
		if err != nil {
			panic(err)
		}
		_, err = db.Sqlx().Queryx("delete from users where chat_id = $1;", testUser.ChatID)
		if err != nil {
			panic(err)
		}
		_, err = db.Sqlx().Queryx("delete from educations where user_id = $1 and book_id = $2;", testUser.ID, testBook.ID)
		if err != nil {
			panic(err)
		}
	}

	bookService := model.NewBookService(gorm.NewBookRepository())
	userService := model.NewUserService(gorm.NewUserRepository())

	testUser, err = userService.Create(testUser)
	if err != nil {
		failFunc(t)
		logger.Get().Error("[TestEducationRepository] Error while create user: %v", err)
		failFunc(t)
		return
	}

	testBook, err = bookService.Create(testBook)
	if err != nil {
		failFunc(t)
		logger.Get().Error("[TestEducationRepository] Error while create book: %v", err)
		failFunc(t)
		return
	}

	repo, err := NewEducationRepository(testUser, testBook)
	if err != nil {
		failFunc(t)
		logger.Get().Error("[TestEducationRepository] Error while create education repository: %v", err)
		failFunc(t)
	}

	educationService := model.NewEducationService(repo)

	t.Run("CreateRelation", func(t *testing.T) {
		if err = educationService.CreateRelation(); err != nil {
			logger.Get().Error("[TestEducationRepository] Error while create relation: %v", err)
			failFunc(t)
		}
	})

	t.Run("GetUnknownWords", func(t *testing.T) {
		_, err = educationService.GetUnknownWords()
		if err != nil {
			logger.Get().Error("[TestEducationRepository] Error while getting unknown words: %v", err)
			failFunc(t)
		}
	})

	t.Run("GetStatistic", func(t *testing.T) {
		processed, err := educationService.GetStatistic()
		if err != nil {
			logger.Get().Error("[TestEducationRepository] Error while getting statistic: %v", err)
			failFunc(t)
		}

		if processed != 0.00 {
			failFunc(t)
		}
	})

	t.Run("SetPoll", func(t *testing.T) {
		err = educationService.SetPoll(1233, 0, 123)
		if err != nil {
			logger.Get().Error("[TestEducationRepository] Error while getting statistic: %v", err)
			failFunc(t)
		}
	})

	t.Run("GetPoll", func(t *testing.T) {
		pollID, err := educationService.GetPoll()
		if err != nil {
			logger.Get().Error("[TestEducationRepository] Error while getting statistic: %v", err)
			failFunc(t)
		}

		if pollID != 1233 {
			failFunc(t)
		}
	})

	clearFunc()
}

func TestEducationRepositoryRealData(t *testing.T) {
	config.NewConfig()
	logger.StartLogger()

	var err error
	user := model.User{
		ID:     2,
		ChatID: 417517295,
		Phone:  "+380664780362",
	}

	book := model.Book{
		ID:   2,
		Name: "Nigel Poulton - The Kubernetes Book (2020).pdf",
	}

	failFunc := func(t *testing.T) {
		t.FailNow()
	}

	repo, err := NewEducationRepository(user, book)
	if err != nil {
		failFunc(t)
		logger.Get().Error("[TestEducationRepository] Error while create education repository: %v", err)
		failFunc(t)
	}

	educationService := model.NewEducationService(repo)

	t.Run("GetUnknownWords", func(t *testing.T) {
		_, err = educationService.GetUnknownWords()
		if err != nil {
			logger.Get().Error("[TestEducationRepository] Error while getting unknown words: %v", err)
			failFunc(t)
		}
	})

	t.Run("GetStatistic", func(t *testing.T) {
		processed, err := educationService.GetStatistic()
		if err != nil {
			logger.Get().Error("[TestEducationRepository] Error while getting statistic: %v", err)
			failFunc(t)
		}

		if processed != 0.00 {
			failFunc(t)
		}
	})

	t.Run("SetPoll", func(t *testing.T) {
		err = educationService.SetPoll(1233, 0, 123)
		if err != nil {
			logger.Get().Error("[TestEducationRepository] Error while getting statistic: %v", err)
			failFunc(t)
		}
	})

	t.Run("GetPoll", func(t *testing.T) {
		pollID, err := educationService.GetPoll()
		if err != nil {
			logger.Get().Error("[TestEducationRepository] Error while getting statistic: %v", err)
			failFunc(t)
		}

		if pollID != 1233 {
			failFunc(t)
		}
	})

}
