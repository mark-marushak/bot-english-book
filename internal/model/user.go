package model

import (
	"gorm.io/gorm"
	"os"
)

const (
	USER_BLOCKED = "block"
	USER_NEW     = "new"
	USER_STUDY   = "study"
)

type User struct {
	gorm.Model
	ChatID    int64  `gorm:"type:bigint"`
	Phone     string `gorm:"type:varchar(50)"`
	Email     string `gorm:"type:varchar(255)"`
	FirstName string `gorm:"type:varchar(255)"`
	Status    string `gorm:"type:varchar(50)"`
	PollID    int
	BookID    uint
	Book      Book
}

type UserService interface {
	Create(user User) error
	Update(user User) error
	Get(user User) (User, error)
	GetKnowingWords(limit, offset int) ([]Word, error)
	UploadBook(file os.File) error
}

type UserRepository interface {
	Create(user User) error
	Update(user User) error
	Get(user User) (User, error)
	GetKnowingWords(limit, offset int) ([]Word, error)
	UploadBook(file os.File) error
}

type userService struct {
	User
	repo UserRepository
}

func NewUserService(repository UserRepository) UserService {
	return &userService{
		repo: repository,
	}
}

func (u userService) Create(user User) error {
	return u.repo.Create(user)
}

func (u userService) Update(user User) error {
	return u.repo.Update(user)
}

func (u userService) GetKnowingWords(limit, offset int) ([]Word, error) {
	return u.repo.GetKnowingWords(limit, offset)
}

func (u userService) UploadBook(file os.File) error {
	return u.repo.UploadBook(file)
}

func (u userService) Get(user User) (User, error) {
	return u.repo.Get(user)
}
