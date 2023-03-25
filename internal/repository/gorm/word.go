package gorm

import (
	translate "cloud.google.com/go/translate/apiv3"
	"context"
	gt "github.com/bas24/googletranslatefree"
	"github.com/mark-marushak/bot-english-book/config"
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/logger"
	"github.com/mark-marushak/bot-english-book/storage"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
	"sync"
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
	text, err := gt.Translate(word.Text, language.English.String(), language.Ukrainian.String())
	if err != nil {
		return nil, err
	}
	word.Text = text

	//var GA_FILE, PROJECT_ID string
	//if err := config.Get().Unmarshal("google.credential-file", &GA_FILE); err != nil {
	//	logger.Get().Error("Error while parsing google api-key %v", err)
	//	return nil, err
	//}
	//
	//if err := config.Get().Unmarshal("google.project-id", &PROJECT_ID); err != nil {
	//	logger.Get().Error("Error while parsing google api-key %v", err)
	//	return nil, err
	//}
	//
	//ctx := context.Background()
	//c, err := translate.NewTranslationClient(ctx, option.WithCredentialsFile(storage.GetGoogleCredsFile(GA_FILE)))
	//if err != nil {
	//	logger.Get().Error("error while creating new translation client: %v", err)
	//	return nil, err
	//}
	//defer c.Close()
	//
	//req := &translatepb.TranslateTextRequest{
	//	Contents:           []string{word.Text},
	//	MimeType:           "text/plain",
	//	SourceLanguageCode: language.AmericanEnglish.String(),
	//	TargetLanguageCode: language.Ukrainian.String(),
	//	Parent:             "projects/" + PROJECT_ID,
	//	Model:              "",
	//	GlossaryConfig:     nil,
	//	Labels:             nil,
	//}
	//
	//resp, err := c.TranslateText(ctx, req)
	//if err != nil {
	//	logger.Get().Error("error while requesting: %v", err)
	//	return nil, err
	//}
	//
	//translations := resp.GetTranslations()
	//
	//if translations == nil {
	//	return nil, fmt.Errorf("translation is not found")
	//}
	//
	//for i := 0; i < len(translations); i++ {
	//	word.Text = translations[i].GetTranslatedText()
	//}
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
