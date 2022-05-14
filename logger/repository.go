package logger

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type LoggerRepository interface {
	Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{})
	//With(field string, value interface{}) LoggerRepository
	//WithError(err error) LoggerRepository
	//Info(context.Context, string)
	//Debug(context.Context, string)
	//Warn(context.Context, string)
	//Error(context.Context, string)
}
