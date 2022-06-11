package internal

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-marushak/bot-english-book/config"
	"github.com/mark-marushak/bot-english-book/logger"
	"github.com/mark-marushak/bot-english-book/pkg/telegram"
)

var botService *BotService

type BotService struct {
	done        chan struct{}
	route       []telegram.RouteService
	telegramBot *tgbotapi.BotAPI
}

func GetBot() *BotService {
	if botService == nil {
		botService = &BotService{}
	}
	return botService
}

func (b BotService) Status() <-chan struct{} {
	return b.done
}

func (b *BotService) Stop() {
	b.done <- struct{}{}
}

func (bs *BotService) Start() {
	var (
		botAPI   string
		botDebug bool
	)

	if err := config.Get().Unmarshal("telegram.bot-api", &botAPI); err != nil {
		logger.Get().Error("error while getting bot-api parameter: %v", err)
	}

	if err := config.Get().Unmarshal("telegram.debug", &botDebug); err != nil {
		logger.Get().Error("error while getting bot-debug parameter: %v", err)
	}

	bot, err := tgbotapi.NewBotAPI(botAPI)
	if err != nil {
		logger.Get().Error("error connecting to telegram bot: %v", err)
	}

	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	bs.telegramBot = bot
	bs.SetRoute()
	logger.Get().Info("Tracking updates is started")
	for update := range updates {
		err = bs.FindRoute(*bot, update)
		if err != nil {
			logger.Get().Error("[Bot]: error while matching route %v", err)
		}

	}
}
