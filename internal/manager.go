package internal

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/internal/repository"
	"os/user"
	"time"
)

// Manager /*
/*
The Manager must control activity of client
If client don't answer long time (10 minutes), Manager will notification
If client forget about lesson, Manager will send notification to start lesson
If any updates became, Manger will say about them to the client through the bot
*/
type Manager struct {
}

func GetManager() *Manager {
	return &Manager{}
}

func (Manager) lookAfterActivity() {
	db.DB().Find(&user.User{})
}

func (m Manager) PrepareBook(id uint) error {
	return repository.CreateAssociation(id)
}

func (m Manager) Start(done chan struct{}) {
	for {
		var books []model.Book
		db.DB().Model(&model.Book{}).Where("status = ?", model.BOOK_UPLOAD).Find(&books)
		if len(books) > 0 {
			for i := 0; i < len(books); i++ {
				// create relationship many2many
				err := m.PrepareBook(books[i].ID)
				if err != nil {
					continue
				}

				books[i].Status = model.BOOK_COMPLETE
				err = repository.NewBookRepository().Update(books[i])
				if err != nil {
					continue
				}

				var users []model.User
				db.DB().Model(&model.User{}).Where("book_id = ?", books[i].ID).Find(&users)

				msg := func(chatID int64) tgbotapi.Chattable {
					return tgbotapi.NewMessage(chatID, "Твоя книжка вже готова, можна починати урок")
				}

				// notify all subscribers
				for i := 0; i < len(users); i++ {
					GetBot().telegramBot.Send(msg(users[i].ChatID))
				}
			}
		}

		time.Sleep(time.Second * 15)
	}
}
