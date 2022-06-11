package sqlx

import (
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/logger"
	"github.com/pemistahl/lingua-go"
	"strings"
)

type languageRepository struct {
	langTable map[string]model.Language
	lang      model.Language
}

func NewLanguageRepository() model.LanguageRepository {
	repo := languageRepository{
		make(map[string]model.Language),
		model.Language{},
	}

	rows, err := db.Sqlx().Queryx("select * from languages")
	if err != nil {
		logger.Get().Error("Error while getting from languages")
	}

	var lang model.Language
	for rows.Next() {
		err = rows.StructScan(&lang)
		if err != nil {
			logger.Get().Error("Error while scanning lanugage: %v", err)
			panic(err)
		}

		repo.cacheResult(lang)
	}

	return &repo
}

func (l languageRepository) GetName(language model.Language) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (l languageRepository) GetCode(language model.Language) (string, error) {
	//TODO implement me
	panic("implement me")
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
	detected := strings.ToLower(language.String())

	if lang := l.searchCache(detected); lang.ID > 0 {
		return lang, nil
	}

	if l.lang.ID == 0 {
		return model.Language{}, model.LanguageNotFoundErr
	}

	return l.lang, nil
}

func (l *languageRepository) cacheResult(lang model.Language) {
	if _, exist := l.langTable[lang.Name]; !exist {
		l.langTable[lang.Name] = lang
	}
}

func (l *languageRepository) searchCache(detected string) model.Language {
	if found, exist := l.langTable[strings.ToLower(detected)]; exist {
		return found
	}

	return model.Language{}
}
