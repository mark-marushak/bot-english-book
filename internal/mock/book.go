package repository

import (
	"github.com/mark-marushak/bot-english-book/internal/model"
	"gorm.io/gorm"
	"os"
)

var books = []model.Book{
	model.Book{
		Model:      gorm.Model{},
		Name:       "Frist Book",
		Complexity: 0,
		Path:       "",
	},
	model.Book{
		Model:      gorm.Model{},
		Name:       "Second Book",
		Complexity: 0,
		Path:       "",
	},
}

type bookRepository struct {
}

func NewBookRepository() *bookRepository {
	return &bookRepository{}
}

func (b bookRepository) FindAll() ([]model.Book, error) {
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

func (b bookRepository) FindByName(name string) (*model.Book, error) {
	for i := 0; i < len(books); i++ {
		if books[i].Name == name {
			return &books[i], nil
		}
	}

	return nil, model.ErrBookNotFound
}
