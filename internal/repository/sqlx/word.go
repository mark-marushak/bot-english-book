package sqlx

import (
	"code.sajari.com/docconv"
	"fmt"
	"github.com/ernestas-poskus/syllables"
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/logger"
	"strings"
	"sync"
	"time"
	"unicode"
)

var wordRepositoryInstance *wordRepository

type wordRepository struct {
	cachedWords *sync.Map
}

func NewWordRepository() *wordRepository {
	if wordRepositoryInstance == nil {
		wordRepositoryInstance = &wordRepository{
			new(sync.Map),
		}
		rows, err := db.Sqlx().Queryx("select * from words")
		if err != nil {
			panic(err)
		}

		var word model.Word
		for rows.Next() {
			err = rows.StructScan(&word)
			if err != nil {
				logger.Get().Error("Error NewWordRepository: %v", err)
				panic(err)
			}

			wordRepositoryInstance.cachedWords.Store(word.Text, word.ID)
		}
	}

	return wordRepositoryInstance
}

func (w wordRepository) GetTranslations(word model.Word) ([]model.Word, error) {
	//TODO implement me
	panic("implement me")
}

func (w wordRepository) GetTranslate(word model.Word) (*model.Word, error) {
	//TODO implement me
	panic("implement me")
}

func (w wordRepository) GetSynonyms(word model.Word) ([]model.Word, error) {
	//TODO implement me
	panic("implement me")
}

func (w wordRepository) Create(word model.Word) (model.Word, error) {
	var id uint
	err := db.Sqlx().QueryRow(`insert into words (text, complexity, language_id, created_at, updated_at) values ($1, $2, $3, $4, $5) RETURNING id`,
		word.Text,
		word.Complexity,
		word.Language.ID,
		time.Now(),
		time.Now()).Scan(&id)

	if err != nil {
		logger.Get().Error("Error while create word: %v", err)
		return model.Word{}, err
	}

	if err == nil {
		word.ID = id
	}

	return word, err
}

func (w wordRepository) Get(word model.Word) (model.Word, error) {
	sql := "select * from words where 1=1"
	if word.ID > 0 {
		sql += fmt.Sprintf(" AND id = %d", word.ID)
	}

	if len(word.Text) > 0 {
		sql += fmt.Sprintf(" AND text = '%s'", word.Text)
	}

	if word.Complexity > 0 {
		sql += fmt.Sprintf(" AND complexity = %d", word.Complexity)
	}

	rows, err := db.Sqlx().Queryx(sql)
	if err != nil {
		logger.Get().Error("Error while getting word: %v", err)
		return model.Word{}, err
	}

	for rows.Next() {
		if err = rows.StructScan(&word); err != nil {
			logger.Get().Error("Error while scan struct word: %v", err)
			return model.Word{}, err
		}
	}

	return word, err
}

func (w wordRepository) Update(word model.Word) (model.Word, error) {
	if tx := db.Gorm().Updates(&word); tx.Error != nil {
		logger.Get().Error("Error while updating word: %v", tx.Error)
		return model.Word{}, tx.Error
	}

	return word, nil
}

func WordGen(words map[string]int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for word, _ := range words {
			select {
			case out <- word:
			}
		}
	}()

	return out
}

func WordFabric(repo model.LanguageService, words <-chan string) <-chan model.Word {
	out := make(chan model.Word)
	go func() {
		defer close(out)
		for word := range words {
			lang, err := repo.DetectLanguage(word)
			if err != nil {
				continue
			}

			created := model.Word{
				Text:       word,
				Complexity: syllables.CountSyllables([]byte(word)),
				Language:   lang,
			}

			out <- created
		}
	}()

	return out
}

func WordTake(repo *wordRepository, wordChan <-chan model.Word) <-chan model.Word {

	wordComplete := make(chan model.Word)
	go func() {
		defer close(wordComplete)
		for word := range wordChan {
			val, _ := repo.cachedWords.Load(word.Text)
			switch v := val.(type) {
			case uint:
				word.ID = v
				wordComplete <- word
				continue
			}

			created, err := repo.Create(word)

			if err != nil {
				return
			}

			repo.cachedWords.Store(created.Text, created.ID)
			wordComplete <- created
		}
	}()

	return wordComplete
}

func CreateAssociation(bookID uint) error {
	repoBook := model.NewBookService(NewBookRepository())
	repoWord := NewWordRepository()
	book, err := repoBook.Get(model.Book{ID: bookID})
	if err != nil {
		logger.Get().Error("err while get book %v", err)
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
		lang, err := languageRepo.DetectLanguage(words[i])
		if err != nil {
			continue
		}

		unique[strings.ToLower(words[i])] = int(lang.ID)
	}

	var b strings.Builder
	var insertValues = make([]string, 0, len(unique))
	b.WriteString("insert into book_words (book_id, word_id) values ")
	for word := range WordTake(repoWord, WordFabric(languageRepo, WordGen(unique))) {
		insertValues = append(insertValues, fmt.Sprintf("(%d, %d)", book.ID, word.ID))
	}

	if len(insertValues) <= 0 {
		return fmt.Errorf("no one words weren't added")
	}

	b.WriteString(strings.Join(insertValues, ", "))
	b.WriteString(" on conflict do nothing")

	db.Sqlx().MustExec(b.String())

	var bw []uint
	err = db.Sqlx().Select(&bw, "select word_id from book_words where book_id = $1", book.ID)
	if err != nil {
		logger.Get().Error("%v", err)
	}

	if len(unique) != len(bw) {
		return fmt.Errorf("assertion words to book not finish unique: %d and book.Words %d", len(unique), len(book.Words))
	}

	return nil
}
