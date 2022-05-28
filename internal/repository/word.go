package repository

import (
	translate "cloud.google.com/go/translate/apiv3"
	"code.sajari.com/docconv"
	"context"
	"errors"
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
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
	"unicode"
)

var ExceptionUnsupportedRelationsError = errors.New("unsupported relations: Books")

type wordRepository struct {
	translationClient *translate.TranslationClient
}

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
	result := db.DB().Create(&word)
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

func WordFabric(words []string) <-chan model.Word {
	wordChan := make(chan model.Word)
	languageRepo := model.NewLanguageService(NewLanguageRepository())

	go func() {
		defer close(wordChan)

		var word string
		for i := 0; i < len(words); i++ {
			word = strings.ToLower(words[i])
			lang, err := languageRepo.DetectLanguage(word)
			if err != nil {
				logger.Get().Error("Detect Language Error: %v", err)
				continue
			}

			created := model.Word{
				Text:       word,
				Frequency:  0,
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
			got, _ := repo.Get(word)
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
		log.Println(err)
		return err
	}

	res, err := docconv.ConvertPath(book.Path)
	if err != nil {
		log.Println(err)
		return err
	}

	words := strings.FieldsFunc(res.Body, func(r rune) bool {
		if unicode.IsLetter(r) {
			return false
		}
		return true
	})

	tx := db.DB().Session(&gorm.Session{PrepareStmt: true})
	for word := range WordTake(WordFabric(words)) {
		time.Sleep(time.Nanosecond * 100)

		result := tx.Exec("INSERT INTO book_words (book_id, word_id) VALUES (?, ?);", book.ID, word.ID)
		if result.Error != nil {
			logger.Get().Error("Insert error: %v", err)
		}
	}
	tx.Commit()

	unique := make(map[string]bool, len(words))
	for i := 0; i < len(words); i++ {
		word := strings.ToLower(words[i])
		unique[word] = true
	}

	db.DB().Preload("Words").Find(&book)
	if len(unique) != len(book.Words) {
		return fmt.Errorf("assertion words to book not finish")
	}

	return nil
}
