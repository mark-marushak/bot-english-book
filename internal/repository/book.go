package repository

import (
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"os"
)

type bookRepository struct {
}

func NewBookRepository() *bookRepository {
	return &bookRepository{}
}

func (b bookRepository) FindAll() ([]model.Book, error) {
	var books []model.Book
	db.DB().Find(&books)
	return books, nil
}

func (b bookRepository) Create(file os.File) error {
	//TODO implement me
	panic("implement me")
}

func (b bookRepository) CalcWord() (int64, error) {
	//TODO implement me
	panic("implement me")
}
