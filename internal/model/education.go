package model

import (
	"gorm.io/gorm"
	"time"
)

type Education struct {
	UserID        uint           `db:"user_id" gorm:"primaryKey"`
	BookID        uint           `db:"book_id" gorm:"primaryKey"`
	PollID        int            `db:"poll_id"`
	CorrectOption int            `db:"correct_option"`
	WordID        uint           `db:"word_id"`
	Processed     float32        `db:"processed"`
	CreatedAt     time.Time      `db:"created_at"`
	UpdatedAt     time.Time      `db:"updated_at"`
	DeletedAt     gorm.DeletedAt `db:"deleted_at" gorm:"index"`
}

type educationService struct {
	repo EducationRepository
}

type EducationRepository interface {
	Get(userID uint) (Education, error)
	Update(userID, bookID uint) error
	CreateRelation() error
	GetUnknownWords() ([]uint, error)
	GetStatistic() (float32, error)
	SetPoll(pollID int, correctOption int, wordID uint) error
	GetPoll() (int, error)
}

type EducationService interface {
	Get(userID uint) (Education, error)
	Update(userID, bookID uint) error
	CreateRelation() error
	GetUnknownWords() ([]uint, error)
	GetStatistic() (float32, error)
	SetPoll(pollID int, correctOption int, wordID uint) error
	GetPoll() (int, error)
}

func NewEducationService(repo EducationRepository) EducationService {
	return &educationService{repo: repo}
}

func (e educationService) CreateRelation() error {
	return e.repo.CreateRelation()
}

func (e educationService) GetUnknownWords() ([]uint, error) {
	return e.repo.GetUnknownWords()
}

func (e educationService) GetStatistic() (float32, error) {
	return e.repo.GetStatistic()
}

func (e educationService) SetPoll(pollID int, correctOption int, wordID uint) error {
	return e.repo.SetPoll(pollID, correctOption, wordID)
}

func (e educationService) GetPoll() (int, error) {
	return e.repo.GetPoll()
}

func (e educationService) Get(userID uint) (Education, error) {
	return e.repo.Get(userID)
}

func (e educationService) Update(userID, bookID uint) error {
	return e.repo.Update(userID, bookID)
}
