package main

import (
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/mark-marushak/bot-english-book/internal"
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/logger"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
)

func main() {
	startLogger()

	db.PrepareTable()

	go handleSystemSignal()
	//go internal.NewManager().Start()
	internal.GetBot().Start()
}

func handleSystemSignal() {
	systemSignal := make(chan os.Signal)
	signal.Notify(systemSignal, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP)

	<-systemSignal

	internal.GetBot().Stop()

	os.Exit(0)
}

func startLogger() {
	_, path, _, _ := runtime.Caller(0)
	path = filepath.Dir(path)

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		path + "error.log",
		"stderr",
	}

	zapLogger, _ := cfg.Build()
	logger.NewLoggerService(zapadapter.NewLogger(zapLogger))
}
