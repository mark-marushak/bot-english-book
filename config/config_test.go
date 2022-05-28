package config

import (
	"github.com/mark-marushak/bot-english-book/logger"
	"net/url"
	"testing"
)

func TestRequestTelegramBot(t *testing.T) {
	logger.StartLogger()
	NewConfig()
	_, err := RequestTelegramBot("getFile", url.Values{"file_id": {"123"}})
	if err == TokenNotFoundError {
		t.Fail()
	}
}
