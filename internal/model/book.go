package model

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var ErrBookNotFound = errors.New("Book isn't found ")

const (
	BOOK_UPLOAD   = "upload"
	BOOK_READ     = "read"
	BOOK_COMPLETE = "complete"
)

type Book struct {
	gorm.Model
	ID         uint `gorm:"primaryKey;index:,unique"`
	MessageID  int
	Name       string  `gorm:"type: varchar(500);index:,unique"`
	Complexity float32 `gorm:"type: decimal(2,2)"`
	Path       string  `gorm:"type: varchar(500);"`
	UserID     uint
	Status     string `gorm:"type: varchar(10)"`
	Words      []Word `gorm:"many2many:book_words;foreignKey:ID;joinForeignKey:BookID;References:ID;joinReferences:WordID"`
}

type bookService struct {
	repo BookRepository
}

type BookRepository interface {
	FindAll() ([]Book, error)
	Create(file Book) (Book, error)
	Update(file Book) error
	Get(book Book) (Book, error)
	CalcWord(file Book) (int64, error)
	FindByName(name string) (Book, error)
}

type BookService interface {
	FindAll() ([]Book, error)
	Create(file Book) (Book, error)
	Update(file Book) error
	Get(book Book) (Book, error)
	CalcWord(file Book) (int64, error)
	FindByName(name string) (Book, error)
}

func NewBookService(repository BookRepository) BookService {
	return &bookService{
		repo: repository,
	}
}

func (b bookService) FindAll() (result []Book, err error) {
	result, err = b.repo.FindAll()
	if err != nil {
		err = fmt.Errorf("Error while find all books: %v", err)
	}
	return
}

func (b bookService) Create(file Book) (Book, error) {
	return b.repo.Create(file)
}

func (b bookService) CalcWord(file Book) (int64, error) {
	return b.repo.CalcWord(file)
}

func (b bookService) FindByName(name string) (Book, error) {
	return b.repo.FindByName(name)
}

func (b bookService) Get(book Book) (Book, error) {
	return b.repo.Get(book)
}

func (b bookService) Update(file Book) error {
	return b.repo.Update(file)
}
