package model

import (
	"gorm.io/gorm"
	"os"
	"os/user"
)

const (
	BLOCKED = "block"
	NEW     = "new"
	STUDY   = "study"
)

type User struct {
	gorm.Model
	ChatID    int64  `gorm: "type:bigint"`
	Phone     string `gorm "type: varchar(50)"`
	FirstName string `gorm: "type: varchar(255)"`
	Books     []Book `gorm: "many2many:user_books;foreignKey:BookID;References:UserID"`
	BookID    uint   `gorm: "index:,unique"`
	Status    string `gorm: "type: varchar(50)"`
}

type UserService interface {
	Create(user User) error
	Update(user user.User) error
	GetKnowingWords(limit, offset int) ([]Word, error)
	UploadBook(file os.File) error
}

type UserRepository interface {
	Create(user User) error
	Update(user user.User) error
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

func (u userService) Update(user user.User) error {
	return u.repo.Update(user)
}

func (u userService) GetKnowingWords(limit, offset int) ([]Word, error) {
	return u.repo.GetKnowingWords(limit, offset)
}

func (u userService) UploadBook(file os.File) error {
	return u.repo.UploadBook(file)
}
