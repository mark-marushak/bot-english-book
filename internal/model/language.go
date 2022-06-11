package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

var LanguageNotDetectedErr = errors.New("language wasn't detected")
var LanguageNotFoundErr = errors.New("language wasn't found in database")

type Language struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" db:"deleted_at"`
	Name      string         `gorm:"type:varchar(255)"`
	Code      string         `gorm:"type:varchar(3)"`
}

func SetupLangs() []Language {
	return []Language{
		{Name: "english", Code: "en"},
		{Name: "ukranian", Code: "ua"},
	}
}

type LanguageRepository interface {
	GetName(Language) (string, error)
	GetCode(Language) (string, error)
	DetectLanguage(string) (Language, error)
}

type LanguageService interface {
	GetName(Language) (string, error)
	GetCode(Language) (string, error)
	DetectLanguage(string) (Language, error)
}

type languageService struct {
	Language
	repo LanguageRepository
}

func NewLanguageService(repository LanguageRepository) LanguageService {
	return &languageService{
		repo: repository,
	}
}

func (s languageService) GetName(lang Language) (string, error) {
	return s.repo.GetName(lang)
}

func (s languageService) GetCode(lang Language) (string, error) {
	return s.repo.GetCode(lang)
}

func (s languageService) DetectLanguage(str string) (Language, error) {
	return s.repo.DetectLanguage(str)
}
