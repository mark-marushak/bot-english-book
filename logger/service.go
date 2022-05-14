package logger

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type LoggerService interface {
	Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{})
	Error(string, ...interface{})
	Info(string, ...interface{})
}
