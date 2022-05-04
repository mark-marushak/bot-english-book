package update

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Callback func(callback tgbotapi.CallbackQuery) tgbotapi.MessageConfig
