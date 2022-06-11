package model

import (
	"gorm.io/gorm"
	"time"
)

type Education struct {
	UserID    uint           `db:"user_id" gorm:"foreignKey"`
	BookID    uint           `db:"book_id" gorm:"foreignKey"`
	Processed float32        `db:"processed"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
	DeletedAt gorm.DeletedAt `db:"deleted_at" gorm:"index"`
}
