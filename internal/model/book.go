package model

import (
	"errors"
	"os"

	"gorm.io/gorm"
)

var ErrBookNotFound = errors.New("Book isn't found ")

type Book struct {
	gorm.Model
	Name       string  `gorm: "type: varchar(500)"`
	Complexity float32 `gorm: "type: decimal(2,2)"`
	Path       string  `gorm: "type: varchar(500);"`
	UserID     uint    `gorm: "index:,unique"`
}

type bookService struct {
	repo BookRepository
}

type BookRepository interface {
	FindAll() ([]Book, error)
	Create(file os.File) error
	CalcWord() (int64, error)
	FindByName(name string) (*Book, error)
}

func NewBookService(repository BookRepository) *bookService {
	return &bookService{
		repo: repository,
	}
}

func (b bookService) FindAll() ([]Book, error) {
	return b.repo.FindAll()
}

func (b bookService) Create(file os.File) error {
	return b.repo.Create(file)
}

func (b bookService) CalcWord() (int64, error) {
	return b.repo.CalcWord()
}

func (b bookService) FindByName(name string) (*Book, error) {
	return b.repo.FindByName(name)
}
