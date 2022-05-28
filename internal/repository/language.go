package repository

import (
	"fmt"
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/pemistahl/lingua-go"
)

type languageRepository struct {
}

func NewLanguageRepository() model.LanguageRepository { return &languageRepository{} }

func (l languageRepository) GetName(lang model.Language) (string, error) {
	result := db.DB().Find(&lang)
	if result.Error != nil {
		return "", result.Error
	}

	return lang.Name, nil
}

func (l languageRepository) GetCode(lang model.Language) (string, error) {
	result := db.DB().Find(&lang)
	if result.Error != nil {
		return "", result.Error
	}

	return lang.Code, nil
}

func (l languageRepository) DetectLanguage(s string) (*model.Language, error) {
	languages := []lingua.Language{
		lingua.English,
		lingua.Ukrainian,
	}

	lingua.AllSpokenLanguages()
	detector := lingua.NewLanguageDetectorBuilder().
		FromLanguages(languages...).
		WithMinimumRelativeDistance(0.25).
		Build()

	var lang model.Language
	if language, exists := detector.DetectLanguageOf(s); exists {
		lang.Name = fmt.Sprintf("%s", language)
		db.DB().Find(&lang)
	} else {
		return nil, model.LanguageNotDetectedErr
	}

	if lang.ID == 0 {
		return nil, model.LanguageNotFoundErr
	}

	return &lang, nil
}
