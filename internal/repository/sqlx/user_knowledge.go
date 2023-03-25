package sqlx

import (
	"database/sql"
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/logger"
)

type userKnowledgeRepository struct {
}

func NewUserKnowledgeRepository() model.UserKnowledgeRepository {
	return userKnowledgeRepository{}
}

func (u userKnowledgeRepository) UpdateLearned(knowledge model.UserKnowledge) error {
	_, err := db.Sqlx().Queryx(`update user_knowledges set learned = $1 where user_id = $2 and word_id = $3`,
		knowledge.Learned,
		knowledge.UserID,
		knowledge.WordID)

	if err != nil {
		logger.Get().Error("Error while UpdateLearned: %v", err)
		return err
	}

	return err
}

func (u userKnowledgeRepository) GetLearned(knowledge model.UserKnowledge) (float32, error) {
	knowledge, err := u.Get(knowledge)
	if err != nil {
		logger.Get().Error("Error while GetLearned: %v", err)
	}
	return knowledge.Learned, err
}

func (u userKnowledgeRepository) StoreAttempt(knowledge model.UserKnowledge) error {
	_, err := db.Sqlx().Queryx(`update user_knowledges set attempt = $1 where user_id = $2 and word_id = $3`,
		knowledge.Attempt+1,
		knowledge.UserID,
		knowledge.WordID)

	if err != nil {
		logger.Get().Error("Error while StoreAttempt: %v", err)
		return err
	}

	return err
}

func (u userKnowledgeRepository) StoreSuccess(knowledge model.UserKnowledge) error {
	_, err := db.Sqlx().Exec(`update user_knowledges set success = $1 where user_id = $2 and word_id = $3`,
		knowledge.Success+1,
		knowledge.UserID,
		knowledge.WordID)

	if err != nil {
		logger.Get().Error("Error while StoreSuccess: %v", err)
		return err
	}

	return err
}

func (u userKnowledgeRepository) Create(knowledge model.UserKnowledge) (model.UserKnowledge, error) {
	_, err := db.Sqlx().Queryx(`insert into user_knowledges (user_id, word_id, learned, attempt, success) values ($1, $2, $3, $4, $5)`,
		knowledge.UserID,
		knowledge.WordID,
		knowledge.Learned,
		knowledge.Attempt,
		knowledge.Success)

	if err != nil {
		logger.Get().Error("Error while Create user_knowledge: %v", err)
		return model.UserKnowledge{}, err
	}

	return knowledge, err
}

func (u userKnowledgeRepository) Get(knowledge model.UserKnowledge) (model.UserKnowledge, error) {
	var get model.UserKnowledge
	err := db.Sqlx().Get(&get, `select * from user_knowledges where user_id = $1 and word_id = $2 limit 1;`,
		knowledge.UserID,
		knowledge.WordID)

	if err == sql.ErrNoRows {
		return get, nil
	}

	if err != nil {
		logger.Get().Error("Error while getting user_knowledge: %v", err)
		return get, err
	}

	return get, err
}

func (u userKnowledgeRepository) GetUserKnowledge(userID uint) ([]model.UserKnowledge, error) {
	var knowledges []model.UserKnowledge
	err := db.Sqlx().Select(&knowledges, `select * from user_knowledges where user_id = $1;`, userID)

	if err == sql.ErrNoRows {
		return knowledges, nil
	}

	if err != nil {
		logger.Get().Error("Error while getting all knowleges: %v", err)
		return knowledges, err
	}

	return knowledges, err
}
