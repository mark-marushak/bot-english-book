package logger

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/mark-marushak/bot-english-book/pkg/logger/zapadapter"
	"github.com/mark-marushak/bot-english-book/storage"
	"go.uber.org/zap"
	"os"
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

func StartLogger() {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		storage.GetStorageRoot() + string(os.PathSeparator) + "error.log",
		"stderr",
	}

	zapLogger, _ := cfg.Build()
	NewLoggerService(zapadapter.NewLogger(zapLogger))
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
