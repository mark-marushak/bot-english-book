package update

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Message func(update tgbotapi.Message) tgbotapi.MessageConfig
