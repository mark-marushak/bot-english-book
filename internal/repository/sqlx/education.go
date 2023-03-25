package sqlx

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/logger"
	"github.com/pkg/errors"
	"time"
)

type educationRepository struct {
	book model.Book
	user model.User
}

func NewEducationRepository(user model.User, book model.Book) (model.EducationRepository, error) {
	if book.ID == 0 {
		return nil, errors.New("EducationRepository book id not found ")
	}

	if user.ID == 0 {
		return nil, errors.New("EducationRepository user id not found ")
	}

	return &educationRepository{book, user}, nil
}

func (e educationRepository) CreateRelation() error {
	_, err := db.Sqlx().Queryx("insert into educations (user_id, book_id, poll_id, correct_option, word_id, processed, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6)",
		e.user.ID,
		e.book.ID,
		0,
		0,
		0,
		0.00,
		time.Now(),
		time.Now(),
		nil)

	if err != nil {
		logger.Get().Error("[EducationRepository] Error while creating relation between book and user in table educations: %v", err)
		return err
	}

	return nil
}

func (e educationRepository) GetUnknownWords() (wordIds []uint, err error) {
	var knownIds []string
	err = db.Sqlx().Select(&knownIds, "select word_id from user_knowledges where user_id = $1 and learned < 100.00", e.user.ID)
	if err != nil {
		logger.Get().Error("[EducationRepository] Error while getting data from user_knowledges: %v", err)
		return nil, err
	}

	var query string
	var args []interface{}
	if len(knownIds) > 0 {
		query, args, err = sqlx.In(fmt.Sprintf("select word_id from book_words where book_id = %d and word_id not in (?)", e.book.ID), knownIds)
		if err != nil {
			logger.Get().Error("[EducationRepository] Error while prepare statement: %v", err)
			return nil, err
		}
	} else {
		query = "select word_id from book_words where book_id = $1"
		args = append(args, e.book.ID)
	}

	query = db.Sqlx().Rebind(query)
	if err != nil {
		logger.Get().Error("[EducationRepository] Error while rebind statement: %v", err)
		return nil, err
	}

	err = db.Sqlx().Select(&wordIds, query, args...)
	if err != nil {
		logger.Get().Error("[EducationRepository] Error while getting data from book_words: %v", err)
		return nil, err
	}

	return wordIds, nil
}

func (e educationRepository) GetStatistic() (float32, error) {
	var processed float32
	err := db.Sqlx().Get(&processed, "select processed from educations where user_id = $1", e.user.ID)
	if err != nil {
		logger.Get().Error("[EducationRepository] Error while getting data from user_knowledges: %v", err)
		return 0.00, err
	}

	return processed, nil
}

func (e educationRepository) SetPoll(pollID, correctOption int, wordID uint) error {
	_, err := db.Sqlx().Queryx(`update educations set poll_id=$1, correct_option=$2, word_id=$3 where user_id = $4 and book_id = $5`,
		pollID,
		correctOption,
		wordID,
		e.user.ID,
		e.book.ID)

	if err != nil {
		logger.Get().Error("Error while create word: %v", err)
		return err
	}

	return err
}

func (e educationRepository) GetPoll() (int, error) {
	var pollID int
	err := db.Sqlx().Get(&pollID, "select poll_id from educations where user_id = $1", e.user.ID)
	if err != nil {
		logger.Get().Error("[EducationRepository] Error while getting data from user_knowledges: %v", err)
		return pollID, err
	}

	return pollID, nil
}

func (e educationRepository) Get(userID uint) (model.Education, error) {
	var education model.Education
	err := db.Sqlx().Get(&education, "select * from educations where user_id = $1", userID)
	if err != nil {
		logger.Get().Error("[EducationRepository] Error while getting data from user_knowledges: %v", err)
		return education, err
	}

	return education, nil
}

func (e educationRepository) Update(userID, bookID uint) error {
	_, err := db.Sqlx().Queryx(`update educations set book_id=$1 where user_id = $2`,
		bookID,
		userID)

	if err != nil {
		logger.Get().Error("Error while update education bookID: %v", err)
		return err
	}

	return err
}
