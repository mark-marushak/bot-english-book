package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-marushak/bot-english-book/src/telegram/sender"
	"regexp"
)

var routes = []route{
	newRoute("Message", "choose book", chooseBook),
	newRoute("Message", "upload book", uploadBook),
	newRoute("Message", "study", studyController),
	newRoute("Callback", "answer ([a-zA-Z]+)", answerController),

	newRoute("Callback", "check stats", checkStats),
}

func newRoute(method tgbotapi.Update, pattern string, handler Handler) Route {
	return Route{method, regexp.MustCompile("^" + pattern), handler}
}

type Handler interface {
	setNext(handler Handler)
	Execute(update *tgbotapi.Update) (int64, string)
}

type Route struct {
	method  tgbotapi.Update
	regex   *regexp.Regexp
	handler Handler
}

func Serve(response *sender.SendService, req tgbotapi.Update) {
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(req.Message.Text)
		if len(matches) > 0 {
			response.Send(route.handler.Execute(&req))
		}
	}

	// not found handle here
}
