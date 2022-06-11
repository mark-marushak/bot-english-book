package model

type UserKnowledge struct {
	UserID  uint    `db:"user_id" gorm:"foreignKey"`
	WordID  uint    `db:"word_id" gorm:"foreignKey"`
	Learned float32 `db:"learned"`
	Attempt int     `db:"attempt"`
	Success int     `db:"success"`
}
