package main

import (
	"github.com/mark-marushak/bot-english-book/config"
	"github.com/mark-marushak/bot-english-book/internal"
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/logger"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	logger.StartLogger()
	config.NewConfig()

	if err := db.PrepareTable(); err != nil {
		logger.Get().Error("Error while preparing database to use: %v", err)
		return
	}

	go internal.GetManager().Start()
	internal.GetBot().Start()

}
