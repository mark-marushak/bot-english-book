package sqlx

import (
	"github.com/mark-marushak/bot-english-book/config"
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/logger"
	"testing"
)

func TestBookRepository(t *testing.T) {
	config.NewConfig()
	logger.StartLogger()
	repo := model.NewBookService(NewBookRepository())

	t.Run("Get", func(t *testing.T) {
		_, err := repo.Get(model.Book{
			ID:   2,
			Name: "Software_Development_Patterns_and_Anti_Patterns_Capers_Jones.pdf",
			Path: "/home/sandbox/development/golang-projects/english-new-words/storage/book/Software_Development_Patterns_and_Anti_Patterns_Capers_Jones.pdf",
		})

		if err != nil {
			t.FailNow()
		}
	})

	t.Run("Create", func(t *testing.T) {
		_, err := db.Sqlx().Queryx("DELETE FROM books where name = $1", "test book")
		if err != nil {
			logger.Get().Error("Error while testing create book: %v", err)
			t.FailNow()
		}

		book, err := repo.Create(model.Book{
			Name: "test book",
			Path: "/some/path",
		})

		if err != nil {
			t.FailNow()
		}

		if book.ID == 0 {
			t.Fail()
		}

		_, err = db.Sqlx().Queryx("DELETE FROM books where name = $1", "test book")
		if err != nil {
			logger.Get().Error("Error while testing create book: %v", err)
			t.FailNow()
		}
	})
}
