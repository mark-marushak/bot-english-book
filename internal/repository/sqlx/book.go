package sqlx

import (
	"code.sajari.com/docconv"
	"fmt"
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/logger"
	"time"
)

type GetData interface {
	GetOne()
	GetAll()
	Find()
}

type bookRepository struct {
	getData GetData
}

func NewBookRepository() model.BookRepository {
	return &bookRepository{}
}

func (b bookRepository) FindAll() ([]model.Book, error) {
	var books []model.Book
	result := db.Gorm().Find(&books)
	return books, result.Error
}

func (b bookRepository) Create(book model.Book) (model.Book, error) {
	var id uint
	err := db.Sqlx().QueryRow(`insert into books (name, complexity, path, user_id, status, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		book.Name,
		book.Complexity,
		book.Path,
		book.UserID,
		book.Status,
		time.Now(),
		time.Now()).Scan(&id)

	if err != nil {
		logger.Get().Error("Error while create book: %v", err)
		return model.Book{}, err
	}

	if err == nil {
		book.ID = id
	}

	return book, err
}

func (b bookRepository) CalcWord(file model.Book) (int64, error) {
	tx := db.Gorm().Find(&file)
	path, _ := tx.Get("path")
	pathString := path.(string)
	res, err := docconv.ConvertPath(pathString)
	if err != nil {
		return 0, err
	}

	fmt.Println(res)

	return 0, nil
}

func (b bookRepository) FindByName(name string) (model.Book, error) {
	var book model.Book
	result := db.Gorm().Model(&model.Book{}).Where("name = ?", name).Find(&book)

	return book, result.Error
}

func (b bookRepository) Get(book model.Book) (model.Book, error) {
	sql := "select * from books where 1=1 "
	if book.ID > 0 {
		sql += fmt.Sprintf(" AND id = %d", book.ID)
	}

	if len(book.Name) > 0 {
		sql += fmt.Sprintf(" AND name = '%s'", book.Name)
	}

	if len(book.Path) > 0 {
		sql += fmt.Sprintf(" AND path = '%s'", book.Path)
	}

	rows, err := db.Sqlx().Queryx(sql)
	if err != nil {
		logger.Get().Error("Error while getting book: %v", err)
		return model.Book{}, err
	}

	for rows.Next() {
		if err = rows.StructScan(&book); err != nil {
			logger.Get().Error("Error while scan struct book: %v", err)
			return model.Book{}, err
		}
	}

	return book, err
}

func (b bookRepository) Update(book model.Book) error {
	result := db.Gorm().Save(&book)
	return result.Error
}
