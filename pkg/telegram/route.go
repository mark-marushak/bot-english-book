package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-marushak/bot-english-book/logger"
	"strings"
)

var ErrNotMatched = errors.New("Path not matched ")

var routeMap *RouteMap

type Route struct {
	path    string
	handler *ServiceHandler
}

func NewRoute(path string, handler *ServiceHandler) *Route {
	return &Route{path, handler}
}

type ReplayRoute struct {
	ID      string
	handler *ServiceHandler
}

func NewReplayRoute(id string, handler *ServiceHandler) *ReplayRoute {
	return &ReplayRoute{id, handler}
}

type RouteMap struct {
	group        map[string][]*Route
	routes       []*Route
	replayRoutes []*ReplayRoute
}

func NewRouteMap() *RouteMap {
	routeMap = &RouteMap{
		map[string][]*Route{},
		[]*Route{},
		[]*ReplayRoute{},
	}

	return routeMap
}

func GetRouteMap() *RouteMap {
	if routeMap == nil {
		logger.Get().Error("[TELEGRAM]: getting route map before it was created", "")
		return nil
	}

	return routeMap
}

func (r RouteMap) Match(update tgbotapi.Update) (handler *ServiceHandler, err error) {
	var text string

	if update.CallbackQuery != nil {
		for i := 0; i < len(r.replayRoutes); i++ {
			if r.replayRoutes[i].ID == update.CallbackQuery.Data {
				return r.replayRoutes[i].handler, nil
			}
		}

		r.replayRoutes = []*ReplayRoute{}
		return nil, ErrNotMatched
	}

	if update.Message != nil {
		text = update.Message.Text
		text = strings.ReplaceAll(text, " ", "-")
		text = strings.TrimSpace(text)
		text = strings.ToLower(text)

		for i := 0; i < len(r.routes); i++ {
			if r.routes[i].path == text {
				return r.routes[i].handler, nil
			}
		}
	}

	return nil, ErrNotMatched
}

func (r *RouteMap) AddGroup(name string, routes []*Route) {
	r.group[name] = make([]*Route, len(routes))
	copy(r.group[name], routes)
}

func (r *RouteMap) AddHandler(route *Route) {
	r.routes = addHandler(route, r.routes)
}

func addHandler(route *Route, routes []*Route) []*Route {
	//list := make([]*Route, len(routes)+1)
	//copy(list, routes)
	//list = append(list, route)
	routes = append(routes, route)
	return routes
}

func (r *RouteMap) AddHandlerCallback(route *ReplayRoute) {
	r.replayRoutes = addHandlerCallback(route, r.replayRoutes)
}

func addHandlerCallback(route *ReplayRoute, routes []*ReplayRoute) []*ReplayRoute {
	//list := make([]*ReplayRoute, len(routes)+1)
	//copy(list, routes)
	//list = append(list, route)
	routes = append(routes, route)
	return routes
}
