package logger

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

var logger LoggerService

type loggerService struct {
	repo LoggerRepository
}

func Get() LoggerService {
	if logger == nil {
		panic("Fatal error with getting logger")
	}

	return logger
}

func NewLoggerService(repository LoggerRepository) LoggerService {
	logger = &loggerService{
		repo: repository,
	}
	return logger
}

func (l loggerService) Error(format string, a ...interface{}) {

	l.repo.Log(context.Background(), pgx.LogLevelError, fmt.Sprintf(format, a), map[string]interface{}{})
}

func (l loggerService) Info(format string, a ...interface{}) {

	l.repo.Log(context.Background(), pgx.LogLevelInfo, fmt.Sprintf(format, a), map[string]interface{}{})
}

func (l loggerService) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	l.repo.Log(ctx, level, msg, data)
}
