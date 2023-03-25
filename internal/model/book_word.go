package model

type BookWord struct {
	BookID    uint `gorm:"primaryKey"`
	WordID    uint `gorm:"primaryKey"`
	Frequency int
}
