package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var NotFoundError = errors.New("Route not found")

/*
RouteService implement middleware after message
based on message or text is sent
return appropriate action that send response message or will do some actions

Basically it is router or group router abstraction
*/

type routeService struct {
	repository RouteRepository
}

type RouteRepository interface {
	Analyze() (int64, error)
	Response(int64) error
	SetBot(*tgbotapi.BotAPI)
	SetUpdate(update tgbotapi.Update)
	SetupRoutes() RouteService
}

type RouteService interface {
	Analyze() (int64, error)
	Response(int64) error
	SetBot(*tgbotapi.BotAPI)
	SetUpdate(update tgbotapi.Update)
	SetupRoutes() RouteService
}

func NewRouteService(repository RouteRepository) RouteService {
	return &routeService{
		repository: repository,
	}
}

func (r routeService) Analyze() (int64, error) {
	return r.repository.Analyze()
}

func (r routeService) Response(i int64) error {
	return r.repository.Response(i)
}

func (r routeService) SetBot(bot *tgbotapi.BotAPI) {
	r.repository.SetBot(bot)
}

func (r routeService) SetUpdate(update tgbotapi.Update) {
	r.repository.SetUpdate(update)
}

func (r *routeService) SetupRoutes() RouteService {
	return r.repository.SetupRoutes()
}
