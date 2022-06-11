package gorm

import (
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"os"
)

type userRepository struct{}

func NewUserRepository() model.UserRepository {
	return &userRepository{}
}

func (u userRepository) Create(user model.User) error {
	result := db.Gorm().Create(&user)
	return result.Error
}

func (u userRepository) Update(user model.User) error {
	result := db.Gorm().Save(&user)
	return result.Error
}

func (u userRepository) GetKnowingWords(limit, offset int) ([]model.Word, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) UploadBook(file os.File) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Get(user model.User) (model.User, error) {
	result := db.Gorm().Where(user).Find(&user)
	return user, result.Error
}
