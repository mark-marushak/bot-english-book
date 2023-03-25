package gorm

import (
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/logger"
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

func (u userRepository) GetEducationByUserID(userID uint) (model.Education, error) {
	var education model.Education
	err := db.Sqlx().Get(&education, "select * from educations where user_id = $1 limit 1", userID)
	if err != nil {
		logger.Get().Error("[EducationRepository] Error while getting data from educations: %v", err)
		return education, err
	}

	return education, nil
}
