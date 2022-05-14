package internal

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/mark-marushak/bot-english-book/logger"
)

var botService *BotService

type BotService struct {
	done chan struct{}
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

func (BotService) Start() {
	var (
		config   = koanf.New(".")
		parser   = yaml.Parser()
		botAPI   string
		botDebug bool
	)

	if err := config.Load(file.Provider("config.yml"), parser); err != nil {
		logger.Get().Error("error loading config: %v", err)
	}

	if err := config.Unmarshal("telegram.bot-api", &botAPI); err != nil {
		logger.Get().Error("error while getting bot-api parameter: %v", err)
	}

	if err := config.Unmarshal("telegram.debug", &botDebug); err != nil {
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
	route := GetRoute()

	logger.Get().Info("Tracking updates is started")

	for update := range updates {
		controller, err := route.Match(update)
		if err != nil {
			logger.Get().Error("[Bot]: error while matching route %v", err)
		}

		msg, err := controller.Send(bot, update)
		if err != nil {
			logger.Get().Error("[Bot]: while sending message %v", err)
		}

		logger.Get().Info("[INFO BOT]: sending message %v", msg)
	}
}
