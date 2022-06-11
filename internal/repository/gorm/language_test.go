package gorm

import (
	"github.com/mark-marushak/bot-english-book/config"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/logger"
	"testing"
)

func TestLanguageRepository(t *testing.T) {
	logger.StartLogger()
	config.NewConfig()

	service := model.NewLanguageService(NewLanguageRepository())

	t.Run("DetectLanguage", func(t *testing.T) {
		_, err := service.DetectLanguage("book")
		if err != nil {
			t.FailNow()
		}
	})

	t.Run("GetCode", func(t *testing.T) {
		lang := model.Language{}
		lang.Name = "english"
		code, err := service.GetCode(lang)
		if err != nil {
			t.FailNow()
		}

		if code != "en" {
			t.Fail()
		}
	})

	t.Run("GetName", func(t *testing.T) {
		lang := model.Language{}
		lang.Code = "en"
		name, err := service.GetName(lang)
		if err != nil {
			t.FailNow()
		}

		if name != "english" {
			t.Fail()
		}
	})

}
