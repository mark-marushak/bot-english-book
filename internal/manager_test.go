package internal

import (
	"github.com/mark-marushak/bot-english-book/config"
	"github.com/mark-marushak/bot-english-book/logger"
	"testing"
)

func TestManager(t *testing.T) {
	config.NewConfig()
	logger.StartLogger()

	t.Run("getUploadedBooks", func(t *testing.T) {

	})

	t.Run("prepareBook", func(t *testing.T) {
		GetManager().Start()
	})

	t.Run("changeStatusBook", func(t *testing.T) {

	})

	t.Run("notifyRelatedUsers", func(t *testing.T) {

	})
}
