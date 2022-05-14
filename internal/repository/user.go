package repository

import (
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"os"
	"os/user"
)

type userRepository struct {
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (u userRepository) Create(user model.User) error {
	result := db.DB().Create(&user)
	return result.Error
}

func (u userRepository) Update(user user.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetKnowingWords(limit, offset int) ([]model.Word, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) UploadBook(file os.File) error {
	//TODO implement me
	panic("implement me")
}
