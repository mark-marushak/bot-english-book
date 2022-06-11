package gorm

import (
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/pemistahl/lingua-go"
)

type languageRepository struct {
	langTable map[string]model.Language
	lang      model.Language
}

func NewLanguageRepository() model.LanguageRepository {
	return &languageRepository{
		make(map[string]model.Language),
		model.Language{},
	}
}

func (l languageRepository) GetName(lang model.Language) (string, error) {
	result := db.Gorm().Find(&lang)
	if result.Error != nil {
		return "", result.Error
	}

	return lang.Name, nil
}

func (l languageRepository) GetCode(lang model.Language) (string, error) {
	result := db.Gorm().Find(&lang)
	if result.Error != nil {
		return "", result.Error
	}

	return lang.Code, nil
}

func (l languageRepository) DetectLanguage(s string) (model.Language, error) {
	languages := []lingua.Language{
		lingua.English,
		lingua.Ukrainian,
	}

	lingua.AllSpokenLanguages()
	detector := lingua.NewLanguageDetectorBuilder().
		FromLanguages(languages...).
		WithMinimumRelativeDistance(0.25).
		Build()

	language, exists := detector.DetectLanguageOf(s)
	if !exists {
		return model.Language{}, model.LanguageNotDetectedErr
	}

	if lang := l.searchCache(language.String()); lang.ID > 0 {
		return lang, nil
	}

	l.lang.Name = language.String()
	db.Gorm().Find(&l.lang)
	l.cacheResult(l.lang)

	if l.lang.ID == 0 {
		return model.Language{}, model.LanguageNotFoundErr
	}

	return l.lang, nil
}

func (l *languageRepository) cacheResult(lang model.Language) {
	if _, exist := l.langTable[lang.Name]; exist {
		l.langTable[lang.Name] = lang
	}
}

func (l *languageRepository) searchCache(detected string) model.Language {
	if found, exist := l.langTable[detected]; exist {
		return found
	}

	return model.Language{}
}
