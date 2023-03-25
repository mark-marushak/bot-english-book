package model

import (
	"gorm.io/gorm"
	"time"
)

const (
	USER_BLOCKED = "block"
	USER_NEW     = "new"
	USER_STUDY   = "study"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	ChatID    int64          `gorm:"primaryKey" db:"chat_id"`
	Phone     string         `gorm:"type:varchar(50)"`
	Email     string         `gorm:"type:varchar(255)"`
	FirstName string         `gorm:"type:varchar(255)" db:"first_name"`
	Status    string         `gorm:"type:varchar(50)"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
	DeletedAt gorm.DeletedAt `db:"deleted_at" gorm:"index"`
}

type UserService interface {
	Create(user User) (User, error)
	Update(user User) (User, error)
	Get(user User) (User, error)
	GetEducationByUserID(userID uint) (Education, error)
}

type UserRepository interface {
	Create(user User) (User, error)
	Update(user User) (User, error)
	Get(user User) (User, error)
	GetEducationByUserID(userID uint) (Education, error)
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

func (u userService) Create(user User) (User, error) {
	return u.repo.Create(user)
}

func (u userService) Update(user User) (User, error) {
	return u.repo.Update(user)
}

func (u userService) Get(user User) (User, error) {
	return u.repo.Get(user)
}

func (u userService) GetEducationByUserID(userID uint) (Education, error) {
	return u.repo.GetEducationByUserID(userID)
}
