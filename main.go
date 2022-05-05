package main

import (
	_ "github.com/mark-marushak/bot-english-book/pkg/bot-plugins/catfacts"
	_ "github.com/mark-marushak/bot-english-book/pkg/bot-plugins/catgif"
	_ "github.com/mark-marushak/bot-english-book/pkg/bot-plugins/encoding"
	"github.com/mark-marushak/bot-english-book/pkg/bot/telegram"
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

func main() {
	var (
		config   = koanf.New(".")
		parser   = yaml.Parser()
		botAPI   string
		botDebug bool
	)

	if err := config.Load(file.Provider("config.yml"), parser); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	if err := config.Unmarshal("telegram.bot-api", &botAPI); err != nil {
		log.Fatalf("error while getting bot-api: %v", err)
	}

	if err := config.Unmarshal("telegram.debug", &botDebug); err != nil {
		log.Fatalf("error while getting bot-api: %v", err)
	}

	telegram.Run(botAPI, botDebug)

	//bot, err := tgbotapi.NewBotAPI(botAPI)
	//if err != nil {
	//	log.Fatalf("error connecting to telegram bot: %v", err)
	//}
	//
	//bot.Debug = true
	//
	//updateConfig := tgbotapi.NewUpdate(0)
	//
	//updateConfig.Timeout = 30
	//
	//updates := bot.GetUpdatesChan(updateConfig)
	//
	//for update := range updates {
	//	if update.Message != nil {
	//		// Construct a new message from the given chat ID and containing
	//		// the text that we received.
	//		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	//
	//		// If the message was open, add a copy of our numeric keyboard.
	//		switch update.Message.Text {
	//		case "open":
	//			msg.ReplyMarkup = numericKeyboard
	//
	//		}
	//
	//		// Send the message.
	//		if _, err = bot.Send(msg); err != nil {
	//			panic(err)
	//		}
	//	} else if update.CallbackQuery != nil {
	//		// Respond to the callback query, telling Telegram to show the user
	//		// a message with the data received.
	//		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	//		if _, err := bot.Request(callback); err != nil {
	//			panic(err)
	//		}
	//
	//		// And finally, send a message containing the data received.
	//		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
	//		if _, err := bot.Send(msg); err != nil {
	//			panic(err)
	//		}
	//	}
	//}
}
