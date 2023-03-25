package sqlx

import (
	"fmt"
	"github.com/mark-marushak/bot-english-book/config"
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordRepository(t *testing.T) {
	logger.StartLogger()
	config.NewConfig()
	repoWord := model.NewWordService(NewWordRepository())
	repoLang := model.NewLanguageService(NewLanguageRepository())

	t.Run("Create", func(t *testing.T) {
		text := "doors"
		_, err := db.Sqlx().Queryx("delete from words where text = $1", text)
		if err != nil {
			logger.Get().Error("Test sqlx: create word err: %v", err)
			t.FailNow()
		}

		lang, _ := repoLang.DetectLanguage(text)
		word := model.Word{
			Text:       text,
			Complexity: 2,
			LanguageID: lang.ID,
		}

		word, err = repoWord.Create(word)
		if err != nil {
			fmt.Println(err)
			t.FailNow()
		}

		if word.ID == 0 {
			t.Fail()
		}
	})

	t.Run("Get", func(t *testing.T) {
		cases := []struct {
			text   string
			expect uint
		}{
			{"brother", uint(5)},
			{"peanuts", uint(6)},
			{"doors", uint(7)},
		}

		for i := 0; i < len(cases); i++ {
			word := model.Word{
				ID:   0,
				Text: cases[i].text,
			}

			word, err := repoWord.Get(word)
			if err != nil {
				t.Fail()
			}

			assert.Equal(t, cases[i].expect, word.ID)
		}

	})

	t.Run("CreateAssociation", func(t *testing.T) {
		err := CreateAssociation(2)
		if err != nil {
			logger.Get().Error("%v", err)
			t.Fail()
		}
	})

	t.Run("GetTranslations", func(t *testing.T) {
		word, err := repoWord.GetTranslate(model.Word{Text: "go"})
		if err != nil {
			t.Fail()
		}

		assert.Equal(t, word.Text, "іди")
	})
}
