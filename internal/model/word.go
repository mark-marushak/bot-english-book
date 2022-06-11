package model

import (
	"gorm.io/gorm"
	"time"
)

//var ExceptionUnsupportedRelationsError = errors.New("unsupported relations: Books")

type Word struct {
	ID         uint           `gorm:"primaryKey;index:,unique"`
	CreatedAt  time.Time      `db:"created_at"`
	UpdatedAt  time.Time      `db:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" db:"deleted_at"`
	Text       string         `gorm:"type:varchar(50);index:,unique"`
	Complexity int            `db:"complexity"`
	LanguageID uint           `db:"language_id"`
	Language   Language
	Books      []Book `gorm:"many2many:book_words;foreignKey:ID;joinForeignKey:BookID;References:ID;joinReferences:WordID"`
}

type WordService interface {
	GetTranslations(Word) ([]Word, error)
	GetTranslate(Word) (*Word, error)
	GetSynonyms(Word) ([]Word, error)
	Create(Word) (Word, error)
	Get(Word) (Word, error)
	Update(Word) (Word, error)
}
type WordRepository interface {
	GetTranslations(Word) ([]Word, error)
	GetTranslate(Word) (*Word, error)
	GetSynonyms(Word) ([]Word, error)
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

func (w wordService) GetTranslations(word Word) ([]Word, error) {
	return w.repo.GetTranslations(word)
}

func (w wordService) GetTranslate(word Word) (*Word, error) {
	return w.repo.GetTranslate(word)
}

func (w wordService) GetSynonyms(word Word) ([]Word, error) {
	return w.repo.GetSynonyms(word)
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
