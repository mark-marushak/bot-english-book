package route

import (
	"github.com/mark-marushak/bot-english-book/internal/action"
	"github.com/mark-marushak/bot-english-book/pkg/telegram"
)

type BookRoute struct {
	baseRoute
}

func (b *BookRoute) SetupRoutes() telegram.RouteService {
	b.route = map[string]map[string]telegram.ActionService{
		"messages": {
			action.OpenLibrary: &action.BookLibrary{},
			action.UploadBook:  &action.BookUpload{},
		},
		"documents": {
			"": &action.BookSend{},
		},
	}

	return b
}

func (b BookRoute) find(list, text string) telegram.ActionService {
	if found, ok := b.route[list][text]; ok {
		return telegram.NewAction(
			found,
		)
	}

	return nil
}

func (b *BookRoute) Analyze() (chatID int64, err error) {
	if b.Update.FromChat() == nil {
		return 0, telegram.RouteNotFoundError
	}

	chatID = b.Update.FromChat().ID

	if b.Update.Message != nil {

		if b.Update.Message.Document != nil {
			b.action = b.find("documents", b.Update.Message.Text)
			return
		}

		text := b.Update.Message.Text
		//text = strings.ToLower(text)
		//text = strings.ReplaceAll(text, " ", "-")

		b.action = b.find("messages", text)
		return
	}

	return 0, telegram.RouteNotFoundError
}
