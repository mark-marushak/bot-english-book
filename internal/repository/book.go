package repository

import (
	"code.sajari.com/docconv"
	"fmt"
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
)

type bookRepository struct {
}

func NewBookRepository() model.BookRepository {
	return &bookRepository{}
}

func (b bookRepository) FindAll() ([]model.Book, error) {
	var books []model.Book
	result := db.DB().Find(books)
	return books, result.Error
}

func (b bookRepository) Create(file model.Book) (model.Book, error) {
	result := db.DB().Create(&file)
	return file, result.Error
}

func (b bookRepository) CalcWord(file model.Book) (int64, error) {
	tx := db.DB().Find(&file)
	path, _ := tx.Get("path")
	pathString := path.(string)
	res, err := docconv.ConvertPath(pathString)
	if err != nil {
		return 0, err
	}

	fmt.Println(res)

	return 0, nil
}

func (b bookRepository) FindByName(name string) (*model.Book, error) {
	var book model.Book
	result := db.DB().Model(&model.Book{}).Where("name = ?", name).Find(&book)
	return &book, result.Error
}

func (b bookRepository) Get(book model.Book) (model.Book, error) {
	result := db.DB().Where(book).Find(&book)
	return book, result.Error
}
