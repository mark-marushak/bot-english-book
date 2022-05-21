package internal

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-marushak/bot-english-book/internal/route"
	"github.com/mark-marushak/bot-english-book/pkg/telegram"
)

func (b *BotService) SetRoute() {
	routes := []telegram.RouteService{
		telegram.NewRouteService(&route.UserRoute{}).SetupRoutes(),
	}
	b.route = make([]telegram.RouteService, len(routes))
	b.route = routes
}

func (b *BotService) FindRoute(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	for i := 0; i < len(b.route); i++ {
		b.route[i].SetBot(bot)
		b.route[i].SetUpdate(update)
		chatID, err := b.route[i].Analyze()

		if err == telegram.NotFoundError {
			continue
		}

		return b.route[i].Response(chatID)
	}

	return notFoundResponse(bot, update)
}

func notFoundResponse(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	message := tgbotapi.NewMessage(update.FromChat().ID, "I can't answer on this message")
	// add some keyboard to this error
	_, err := bot.Send(message)

	if err != nil {
		return err
	}

	return fmt.Errorf("route wasn't found. Client get response")
}
