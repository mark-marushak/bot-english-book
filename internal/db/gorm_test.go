package db

import (
	"github.com/mark-marushak/bot-english-book/config"
	"github.com/mark-marushak/bot-english-book/logger"
	"testing"
)

func TestGorm(t *testing.T) {
	config.NewConfig()
	logger.StartLogger()
	t.Run("PrepareTable", func(t *testing.T) {
		err := PrepareTable()
		if err != nil {
			t.Fail()
		}
	})

	t.Run("Gorm", func(t *testing.T) {
		if Gorm().Error != nil {
			t.Fail()
		}
	})
}
