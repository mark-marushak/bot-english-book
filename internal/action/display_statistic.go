package action

import (
	"fmt"
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/internal/repository/gorm"
	"github.com/mark-marushak/bot-english-book/internal/repository/sqlx"
)

type DisplayStatistic struct {
	AdaptorTelegramAction
}

func (d DisplayStatistic) Keyboard(i ...interface{}) interface{} {
	return BackToMainMenu
}

func (d DisplayStatistic) Output(i ...interface{}) (string, error) {
	var user model.User
	user.ChatID = d.GetUpdate().FromChat().ID

	userService := model.NewUserService(gorm.NewUserRepository())
	user, err := userService.Get(user)
	if err != nil {
		return "", err
	}

	education, err := userService.GetEducationByUserID(user.ID)
	if err != nil {
		return "", err
	}

	// count knowledge
	knowledgeService := model.NewUserKnowledgeService(sqlx.NewUserKnowledgeRepository())
	knowledges, err := knowledgeService.GetUserKnowledge(user.ID)
	if err != nil {
		return "", err
	}

	var words []struct {
		BookID    uint `db:"book_id"`
		WordID    uint `db:"word_id"`
		Frequency interface{}
	}
	//difference words in book and in knowledge
	err = db.Sqlx().Select(&words, "select * from book_words where book_id = $1", education.BookID)
	if err != nil {
		return "", err
	}

	// percent of processed
	var count int
	for i := 0; i < len(knowledges); i++ {
		for j := 0; j < len(words); j++ {
			if words[j].WordID == knowledges[i].WordID && knowledges[i].Success >= knowledges[i].Attempt {
				count++
			}
		}
	}

	var progress float32 = float32(count) / float32(len(knowledges))

	return fmt.Sprintf("%s ваша статистико до цієї книжки:\n*знаєте слів - %d\n*залишилось вивчити - %d\n*прогресс - %.2f%%                ",
		user.FirstName,
		count,
		len(words)-count,
		progress,
	), nil
}
