package model

import (
	"gorm.io/gorm"
	"time"
)

type Education struct {
	UserID    uint `db:"user_id" gorm:"primaryKey"`
	User      User
	BookID    uint `db:"book_id" gorm:"primaryKey"`
	Book      Book
	Processed float32        `db:"processed"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
	DeletedAt gorm.DeletedAt `db:"deleted_at" gorm:"index"`
}

type educationService struct {
	repo EducationRepository
}

type EducationRepository interface {
	CreateRelation() error
	GetUnknownWords() ([]uint, error)
	GetStatistic() (float32, error)
}

type EducationService interface {
	CreateRelation() error
	GetUnknownWords() ([]uint, error)
	GetStatistic() (float32, error)
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
