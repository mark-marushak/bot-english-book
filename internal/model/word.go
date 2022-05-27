package model

import (
	"gorm.io/gorm"
)

type Word struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey;index:,unique"`
	Text       string `gorm:"type: varchar(50);index:,unique"`
	Frequency  int
	Complexity int
	LanguageID uint
	Language   Language
}

type WordService interface {
	GetTranslations() []Word
	GetSynonyms() []Word
	Create(Word) (Word, error)
	Get(Word) (Word, error)
	Update(Word) (Word, error)
}
type WordRepository interface {
	GetTranslations() []Word
	GetSynonyms() []Word
	Create(Word) (Word, error)
	Get(Word) (Word, error)
	Update(Word) (Word, error)
}

type wordService struct {
	Word
	repo WordRepository
}

func NewWordService(repository WordRepository) WordService {
	return &wordService{
		repo: repository,
	}
}

func (w wordService) GetTranslations() []Word {
	return w.repo.GetTranslations()
}

func (w wordService) GetSynonyms() []Word {
	return w.repo.GetSynonyms()
}

func (w wordService) Create(word Word) (Word, error) {
	return w.repo.Create(word)
}

func (w wordService) Update(word Word) (Word, error) {
	return w.repo.Update(word)
}

func (w wordService) Get(word Word) (Word, error) {
	return w.repo.Get(word)
}
