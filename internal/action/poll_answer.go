package action

import (
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/internal/repository/gorm"
	"github.com/mark-marushak/bot-english-book/internal/repository/sqlx"
)

type PollAnswer struct {
	AdaptorTelegramAction
}

func (p PollAnswer) Keyboard(i ...interface{}) interface{} {
	return nil
}

func (p PollAnswer) Output(i ...interface{}) (string, error) {
	userService := model.NewUserService(gorm.NewUserRepository())
	user, err := userService.Get(model.User{ChatID: p.GetUpdate().PollAnswer.User.ID})
	if err != nil {
		return "", err
	}

	education, err := userService.GetEducationByUserID(user.ID)
	if err != nil {
		return "", err
	}

	knowledgeService := model.NewUserKnowledgeService(sqlx.NewUserKnowledgeRepository())
	knowledge := model.UserKnowledge{
		UserID:  user.ID,
		WordID:  education.WordID,
		Learned: 0,
		Attempt: 0,
		Success: 0,
	}

	get, err := knowledgeService.Get(knowledge)
	if err != nil {
		return "", err
	}

	if get.UserID <= 0 {
		knowledge, err = knowledgeService.Create(knowledge)
		if err != nil {
			return "", err
		}
	} else {
		knowledge = get
	}

	var output string
	if education.CorrectOption == p.GetUpdate().PollAnswer.OptionIDs[0] {
		err = knowledgeService.StoreSuccess(knowledge)
		if err != nil {
			return "", err
		}

		output = "Виправильно відповіли!"
	} else {
		output = "Відповідь не правильна("
	}

	err = knowledgeService.StoreAttempt(knowledge)
	if err != nil {
		return "", err
	}

	return output, nil
}
