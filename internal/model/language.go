package model

import "gorm.io/gorm"

type Language struct {
	gorm.Model
	Name string `gorm: "type: varchar(255)"`
}

type LanguageRepository interface {
	GetName() string
}

type LanguageService interface {
	GetName() string
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

func (s languageService) GetName() string {
	return s.repo.GetName()
}
