package internal

import (
	"github.com/mark-marushak/bot-english-book/internal/handler"
	"github.com/mark-marushak/bot-english-book/pkg/telegram"
)

const (
	ChooseBook = "choose-book"
)

func GetRoute() *telegram.RouteMap {

	routeMap := telegram.NewRouteMap()
	// todo: meke permission for route map
	// if the group has permission on this route then it pass
	// if the route stay after or before then it can't be pass
	//routeMap.AddGroup(ChooseBook, []*telegram.Route{
	//	telegram.NewRoute("*.pdf", telegram.NewHandler(handler.BookHanlder{})),
	//})

	routeMap.AddHandler(telegram.NewRoute(ChooseBook, telegram.NewHandler(handler.ChooseBookHandler{})))

	return routeMap

}
