package gorm

import (
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
)

type userRepository struct{}

func NewUserRepository() model.UserRepository {
	return &userRepository{}
}

func (u userRepository) Create(user model.User) (model.User, error) {
	result := db.Gorm().Create(&user)
	return user, result.Error
}

func (u userRepository) Update(user model.User) (model.User, error) {
	result := db.Gorm().Save(&user)
	return user, result.Error
}

func (u userRepository) Get(user model.User) (model.User, error) {
	result := db.Gorm().Where(user).Find(&user)
	return user, result.Error
}
