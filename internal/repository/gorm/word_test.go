package gorm

import (
	"fmt"
	"github.com/mark-marushak/bot-english-book/config"
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
		text := "frequency"
		lang, _ := repoLang.DetectLanguage(text)
		word := model.Word{
			Text:       text,
			Complexity: 2,
			Language:   lang,
		}

		word, err := repoWord.Create(word)
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
			{"brother", uint(2276)},
			{"peanuts", uint(2278)},
			{"doors", uint(2280)},
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
		err := CreateAssociation(7)
		if err != nil {
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
