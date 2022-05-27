package repository

import (
	"fmt"
	"github.com/mark-marushak/bot-english-book/config"
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/logger"
	"log"
	"testing"
)

var testBookPath = "/home/sanbox/development/golang-projects/english-new-words/storage/book/difficult-situations-quiz.pdf"

func TestBookRepository(t *testing.T) {
	logger.StartLogger()
	config.NewConfig()
	db.PrepareTable()

	t.Run("Create", func(t *testing.T) {
		repo := model.NewBookService(NewBookRepository())
		_, err := repo.Create(model.Book{
			MessageID:  123,
			Name:       "difficult situations quiz.pdf",
			Complexity: 99.99,
			Path:       testBookPath,
			UserID:     1,
		})

		if err != nil {
			log.Println(err)
			t.Fail()
		}
	})

	t.Run("CalcWords", func(t *testing.T) {
		book := model.Book{}
		book.Name = "difficult situations quiz.pdf"
		db.DB().Where(book).Preload("Words").Find(&book)

		if len(book.Words) <= 0 {
			t.Fail()
		}

		fmt.Println()
	})
}
