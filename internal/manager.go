package internal

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mark-marushak/bot-english-book/internal/db"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/internal/repository/gorm"
	"github.com/mark-marushak/bot-english-book/internal/repository/sqlx"
	"github.com/mark-marushak/bot-english-book/logger"
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

func (m Manager) getUploadedBooks() (books []model.Book, err error) {
	tx := db.Gorm().Model(&model.Book{}).Where("status = ?", model.BOOK_UPLOAD).Find(&books)
	return books, tx.Error
}

func (m Manager) prepareBook(id uint) error {
	logger.Get().Info("Start preparing a book %d", id)
	return sqlx.CreateAssociation(id)
}

func (m Manager) changeStatusBook(book model.Book) error {
	book.Status = model.BOOK_COMPLETE

	err := gorm.NewBookRepository().Update(book)
	if err != nil {
		return err
	}

	return nil
}

func (m Manager) notifyRelatedUsers(book model.Book) (err error) {
	var users []model.User
	tx := db.Gorm().Model(&model.User{}).Where("book_id = ?", book.ID).Find(&users)
	if tx.Error != nil {
		return tx.Error
	}

	msg := func(chatID int64) tgbotapi.Chattable {
		return tgbotapi.NewMessage(chatID, "Твоя книжка вже готова, можна починати урок")
	}

	for i := 0; i < len(users); i++ {
		_, err = GetBot().telegramBot.Send(msg(users[i].ChatID))
		if err != nil {
			return err
		}
	}

	return nil
}

func (m Manager) Start() {
	for {
		books, err := m.getUploadedBooks()
		if err != nil {
			logger.Get().Error("Error while getting uploaded books: %v", err)
			continue
		}

		if len(books) > 0 {
			for i := 0; i < len(books); i++ {
				err = m.prepareBook(books[i].ID)
				if err != nil {
					logger.Get().Error("Error while prepare book: %v", err)
					continue
				}

				err = m.changeStatusBook(books[i])
				if err != nil {
					logger.Get().Error("The book %d was prepared %v", err)
				}

				err = m.notifyRelatedUsers(books[i])
				if err != nil {
					logger.Get().Error("The message wasn't reached the client. Client didn't get notification: %v", err)
				}

			}
		}

		time.Sleep(time.Second * 15)
	}
}
