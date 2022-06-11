package gorm

import (
	translate "cloud.google.com/go/translate/apiv3"
	"code.sajari.com/docconv"
	"context"
	"fmt"
	"github.com/ernestas-poskus/syllables"
	"github.com/mark-marushak/bot-english-book/config"
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/logger"
	"github.com/mark-marushak/bot-english-book/storage"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
	"strings"
	"sync"
	"unicode"
)

var cachedWords *sync.Map

type wordRepository struct{}

func NewWordRepository() model.WordRepository {
	return &wordRepository{}
}

func (w wordRepository) GetTranslations(word model.Word) ([]model.Word, error) {
	var GA_FILE, PROJECT_ID string
	if err := config.Get().Unmarshal("google.credential-file", &GA_FILE); err != nil {
		logger.Get().Error("Error while parsing google api-key %v", err)
		return nil, err
	}

	if err := config.Get().Unmarshal("google.project-id", &PROJECT_ID); err != nil {
		logger.Get().Error("Error while parsing google api-key %v", err)
		return nil, err
	}

	ctx := context.Background()
	c, err := translate.NewTranslationClient(ctx, option.WithCredentialsFile(storage.GetGoogleCredsFile(GA_FILE)))
	if err != nil {
		logger.Get().Error("error while creating new translation client: %v", err)
		return nil, err
	}
	defer c.Close()

	req := &translatepb.TranslateTextRequest{
		Contents:           []string{word.Text},
		MimeType:           "text/plain",
		SourceLanguageCode: language.AmericanEnglish.String(),
		TargetLanguageCode: language.Ukrainian.String(),
		Parent:             "projects/" + PROJECT_ID,
		Model:              "",
		GlossaryConfig:     nil,
		Labels:             nil,
	}

	resp, err := c.TranslateText(ctx, req)
	if err != nil {
		logger.Get().Error("error while requesting: %v", err)
		return nil, err
	}

	logger.Get().Info("Data response: %v", resp)
	return nil, nil
}

func (w wordRepository) GetTranslate(word model.Word) (*model.Word, error) {
	var GA_FILE, PROJECT_ID string
	if err := config.Get().Unmarshal("google.credential-file", &GA_FILE); err != nil {
		logger.Get().Error("Error while parsing google api-key %v", err)
		return nil, err
	}

	if err := config.Get().Unmarshal("google.project-id", &PROJECT_ID); err != nil {
		logger.Get().Error("Error while parsing google api-key %v", err)
		return nil, err
	}

	ctx := context.Background()
	c, err := translate.NewTranslationClient(ctx, option.WithCredentialsFile(storage.GetGoogleCredsFile(GA_FILE)))
	if err != nil {
		logger.Get().Error("error while creating new translation client: %v", err)
		return nil, err
	}
	defer c.Close()

	req := &translatepb.TranslateTextRequest{
		Contents:           []string{word.Text},
		MimeType:           "text/plain",
		SourceLanguageCode: language.AmericanEnglish.String(),
		TargetLanguageCode: language.Ukrainian.String(),
		Parent:             "projects/" + PROJECT_ID,
		Model:              "",
		GlossaryConfig:     nil,
		Labels:             nil,
	}

	resp, err := c.TranslateText(ctx, req)
	if err != nil {
		logger.Get().Error("error while requesting: %v", err)
		return nil, err
	}

	translations := resp.GetTranslations()

	if translations == nil {
		return nil, fmt.Errorf("translation is not found")
	}

	for i := 0; i < len(translations); i++ {
		word.Text = translations[i].GetTranslatedText()
	}
	return &word, nil
}

func (w wordRepository) GetSynonyms(word model.Word) ([]model.Word, error) {
	//TODO implement me
	panic("implement me")
}

func (w wordRepository) Create(word model.Word) (model.Word, error) {
	result := db.Gorm().Create(&word)
	return word, result.Error
}

func (w wordRepository) Get(word model.Word) (model.Word, error) {
	result := db.Gorm().Where(word).Find(&word)
	return word, result.Error
}

func (w wordRepository) Update(word model.Word) (model.Word, error) {
	//TODO implement me
	panic("implement me")
}

func WordGen(words map[string]int) <-chan string {
	out := make(chan string)
	go func() {
		for word, _ := range words {
			select {
			case out <- word:
			}
		}
	}()

	return out
}

func WordFabric(words <-chan string) <-chan model.Word {
	wordChan := make(chan model.Word)
	languageRepo := model.NewLanguageService(NewLanguageRepository())

	go func() {
		defer close(wordChan)

		var word string
		for i := 0; i < len(words); i++ {
			word = strings.ToLower(word)
			lang, err := languageRepo.DetectLanguage(word)
			if err != nil {
				continue
			}

			created := model.Word{
				Text:       word,
				Complexity: syllables.CountSyllables([]byte(word)),
				LanguageID: lang.ID,
			}

			wordChan <- created
		}
	}()

	return wordChan
}

func WordTake(wordChan <-chan model.Word) <-chan model.Word {
	repo := NewWordRepository()
	wordComplete := make(chan model.Word)
	go func() {
		defer close(wordComplete)
		for word := range wordChan {
			var got model.Word
			db.Gorm().Raw("select * from words where text = ?", word.Text).Scan(&got)
			if got.ID > 0 {
				wordComplete <- got
				continue
			}

			created, _ := repo.Create(word)
			if created.ID > 0 {
				wordComplete <- created
			}
		}
	}()

	return wordComplete
}

func CreateAssociation(bookID uint) error {
	repoBook := model.NewBookService(NewBookRepository())
	book, err := repoBook.Get(model.Book{ID: bookID})
	if err != nil {
		return err
	}

	res, err := docconv.ConvertPath(book.Path)
	if err != nil {
		return err
	}

	words := strings.FieldsFunc(res.Body, func(r rune) bool {
		if unicode.IsLetter(r) {
			return false
		}
		return true
	})

	unique := make(map[string]int, len(words))
	languageRepo := model.NewLanguageService(NewLanguageRepository())
	for i := 0; i < len(words); i++ {
		_, err = languageRepo.DetectLanguage(words[i])
		if err != nil {
			continue
		}

		unique[strings.ToLower(words[i])]++
	}

	book.Words = make([]model.Word, 0, len(unique))
	dbSession := db.Gorm().WithContext(context.Background())
	for word := range WordTake(WordFabric(WordGen(unique))) {
		book.Words = append(book.Words, word)
		dbSession.Exec("insert into book_words (book_id, word_id) values (?, ?) on conflict do nothing", book.ID, word.ID)
	}

	db.Gorm().Preload("Words").Find(&book)
	if len(unique)-2 > len(book.Words) {
		return fmt.Errorf("assertion words to book not finish unique: %d and book.Words %d", len(unique), len(book.Words))
	}

	return nil
}
