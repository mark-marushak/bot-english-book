package repository

import (
	"errors"
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
)

var ExceptionUnsupportedRelationsError = errors.New("unsupported relations: Books")

type wordRepository struct{}

func NewWordRepository() model.WordRepository {
	return &wordRepository{}
}

func (w wordRepository) GetTranslations() []model.Word {
	//TODO implement me
	panic("implement me")
}

func (w wordRepository) GetSynonyms() []model.Word {
	//TODO implement me
	panic("implement me")
}

func (w wordRepository) Create(word model.Word) (model.Word, error) {
	result := db.DB().Create(&word)
	//err = db.DB().Model(&model.Book{}).Association("Books").Append([]*model.Word{&word})
	//if err != nil {
	//	return
	//}
	return word, result.Error
}

func (w wordRepository) Get(word model.Word) (model.Word, error) {
	result := db.DB().Where(word).Find(&word)
	return word, result.Error
}

func (w wordRepository) Update(word model.Word) (model.Word, error) {
	//TODO implement me
	panic("implement me")
}
