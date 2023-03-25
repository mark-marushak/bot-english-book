package sqlx

import (
	"github.com/mark-marushak/bot-english-book/config"
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/logger"
	"testing"
)

func TestUserKnowledgeRepository(t *testing.T) {
	config.NewConfig()
	logger.StartLogger()

	var err error
	testUser := model.User{
		ID:        123123123,
		ChatID:    123,
		Phone:     "123123123",
		Email:     "mm@gmail.com",
		FirstName: "test user",
		Status:    "new",
	}

	testWord := model.Word{
		ID:         9999999,
		Text:       "asdfasdfasdfasdfaf",
		LanguageID: 1,
		Complexity: 4,
	}

	failFunc := func(t *testing.T) {
		db.Sqlx().Queryx("delete from user_knowledges where user_id = $1 and word_id = $2", testUser.ID, testWord.ID)
		t.FailNow()
	}

	clearFunc := func() {
		_, err = db.Sqlx().Queryx("delete from user_knowledges where user_id = $1 and word_id = $2;", testUser.ID, testWord.ID)
		if err != nil {
			panic(err)
		}
	}

	service := model.NewUserKnowledgeService(NewUserKnowledgeRepository())

	t.Run("Create", func(t *testing.T) {
		knowlage := model.UserKnowledge{
			UserID:  testUser.ID,
			WordID:  testWord.ID,
			Learned: 0,
			Attempt: 0,
			Success: 0,
		}

		_, err = service.Create(knowlage)
		if err != nil {
			failFunc(t)
		}
	})

	t.Run("Get", func(t *testing.T) {
		knowlage := model.UserKnowledge{
			UserID: testUser.ID,
			WordID: testWord.ID,
		}

		_, err = service.Get(knowlage)
		if err != nil {
			failFunc(t)
		}
	})

	t.Run("UpdateLearned", func(t *testing.T) {
		knowlage := model.UserKnowledge{
			UserID:  testUser.ID,
			WordID:  testWord.ID,
			Learned: 97.97,
		}

		err = service.UpdateLearned(knowlage)
		if err != nil {
			failFunc(t)
		}

		k, _ := service.Get(knowlage)
		if k.Learned != knowlage.Learned {
			failFunc(t)
		}
	})

	t.Run("GetLearned", func(t *testing.T) {
		knowlage := model.UserKnowledge{
			UserID: testUser.ID,
			WordID: testWord.ID,
		}

		_, err = service.GetLearned(knowlage)
		if err != nil {
			failFunc(t)
		}
	})

	t.Run("StoreAttempt", func(t *testing.T) {
		knowlage := model.UserKnowledge{
			UserID:  testUser.ID,
			WordID:  testWord.ID,
			Attempt: 10,
		}

		err = service.StoreAttempt(knowlage)
		if err != nil {
			failFunc(t)
		}

		k, _ := service.Get(knowlage)
		if k.Attempt != 10 {
			failFunc(t)
		}
	})

	t.Run("StoreSuccess", func(t *testing.T) {
		knowlage := model.UserKnowledge{
			UserID:  testUser.ID,
			WordID:  testWord.ID,
			Success: 10,
		}

		err = service.StoreSuccess(knowlage)
		if err != nil {
			failFunc(t)
		}

		k, _ := service.Get(knowlage)
		if k.Success != 10 {
			failFunc(t)
		}
	})

	clearFunc()
}

func TestUserKnowledgeRepositoryRealData(t *testing.T) {
	config.NewConfig()
	logger.StartLogger()

	service := model.NewUserKnowledgeService(NewUserKnowledgeRepository())

	t.Run("GetUserKnowledge", func(t *testing.T) {
		knowledges, err := service.GetUserKnowledge(2)
		if err != nil {
			return
		}

		if len(knowledges) < 6 {
			t.FailNow()
		}

	})
}
