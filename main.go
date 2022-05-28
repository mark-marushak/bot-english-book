package main

import (
	"github.com/mark-marushak/bot-english-book/config"
	"github.com/mark-marushak/bot-english-book/internal"
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/logger"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
)

func main() {
	logger.StartLogger()
	config.NewConfig()

	db.PrepareTable()

	done := make(chan struct{})
	go internal.GetManager().Start(done)
	go handleSystemSignal(done)
	//go internal.NewManager().Start()
	internal.GetBot().Start()
}

func handleSystemSignal(done chan struct{}) {
	systemSignal := make(chan os.Signal)
	signal.Notify(systemSignal, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP)

	<-systemSignal
	close(done)

	internal.GetBot().Stop()

	os.Exit(0)
}

func RootFolder() string {
	_, b, _, ok := runtime.Caller(0)

	if !ok {
		log.Fatal("[ERR]: RootFolder ")
	}

	return filepath.Dir(b)
}
