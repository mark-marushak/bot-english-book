package repository

import (
	"code.sajari.com/docconv"
	"fmt"
	"github.com/ernestas-poskus/syllables"
	"github.com/mark-marushak/bot-english-book/config"
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/logger"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"sync"
	"testing"
	"unicode"
)

func TestWordRepository(t *testing.T) {
	logger.StartLogger()
	config.NewConfig()
	db.PrepareTable()
	repoBook := model.NewBookService(NewBookRepository())
	repoWord := model.NewWordService(NewWordRepository())
	repoLang := model.NewLanguageService(NewLanguageRepository())

	t.Run("Create", func(t *testing.T) {
		text := "frequency"
		lang, _ := repoLang.DetectLanguage(text)
		word := model.Word{
			Text:       text,
			Frequency:  1,
			Complexity: 2,
			Language:   *lang,
		}

		word, err := repoWord.Create(word)
		if err != nil {
			fmt.Println(err)
			t.FailNow()
		}
	})

	t.Run("Get", func(t *testing.T) {
		cases := []struct {
			text   string
			expect uint
		}{
			{"father", uint(165)},
			{"since", uint(173)},
			{"constantly", uint(180)},
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
		book, err := repoBook.FindByName("difficult situations quiz.pdf")
		if err != nil {
			log.Println(err)
			t.FailNow()
		}

		res, err := docconv.ConvertPath(book.Path)
		if err != nil {
			log.Println(err)
			t.FailNow()
		}

		words := strings.FieldsFunc(res.Body, func(r rune) bool {
			if unicode.IsLetter(r) {
				return false
			}
			return true
		})

		wordsMap := make(map[string]int, len(words))
		for i := 0; i < len(words); i++ {
			wordsMap[strings.ToLower(words[i])]++
		}

		var once sync.Once
		cancel := func(done chan bool) {
			once.Do(func() {
				close(done)
			})
		}

		fabrica := func(done chan bool, words map[string]int) <-chan model.Word {
			wordChan := make(chan model.Word)
			languageRepo := model.NewLanguageService(NewLanguageRepository())

			go func() {
				defer close(wordChan)
				for word, count := range words {
					lang, err := languageRepo.DetectLanguage(word)
					if err != nil {
						logger.Get().
							Error("Detect Language Error: %v", err)
						cancel(done)
					}

					select {
					case <-done:
						return
					default:
					}

					wordChan <- model.Word{
						Text:       word,
						Frequency:  count,
						Complexity: syllables.CountSyllables([]byte(word)),
						LanguageID: lang.ID,
					}
				}
			}()

			return wordChan
		}

		recorder := func(done chan bool, wordChan <-chan model.Word) <-chan model.Word {
			wordComplete := make(chan model.Word)
			go func() {
				defer close(wordComplete)
				for word := range wordChan {

					newWord, err := repoWord.Get(word)
					logger.Get().Info("word got: %#v", newWord)
					if err != nil {
						logger.Get().Error("Test Err Getting in pipeline: %v", err)
						newWord, err = repoWord.Create(word)
						if err != nil {
							logger.Get().Error("Test Err Creating in pipeline: %v", err)
						}
					}

					select {
					case <-done:
						return
					case wordComplete <- newWord:
					}
				}
			}()

			return wordComplete
		}

		done := make(chan bool)
		association := make([]model.Word, 0, len(wordsMap))
		pipeline := recorder(done, fabrica(done, wordsMap))
		for word := range pipeline {
			logger.Get().Error("Creaeted word %v", word)
			association = append(association, word)
		}

		if err := db.DB().Model(&book).Association("Words").Append(association); err != nil {
			logger.Get().Error("Append error %v", err)
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
