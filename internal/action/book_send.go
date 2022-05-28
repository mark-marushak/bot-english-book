package action

import (
	"encoding/json"
	"fmt"
	"github.com/mark-marushak/bot-english-book/config"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/internal/repository"
	"github.com/mark-marushak/bot-english-book/logger"
	"github.com/mark-marushak/bot-english-book/storage"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type BookSend struct {
	AdaptorTelegramAction
}

//var textBookSend = `Ok, start with this book. Press button to get your first lesson`
var textBookSend = `Супер, твоя нова книжка додана. 
Почекай повідомлення від мене, 
що книжка готова до навчання
Це зазвичай дуже швидка процедура! 1-2 хвилини`

func (b BookSend) Keyboard(i ...interface{}) interface{} {
	return StartStudyButton
}

func (b BookSend) Output(i ...interface{}) (string, error) {
	var document = b.GetUpdate().Message.Document

	name := document.FileName
	name = strings.ReplaceAll(name, " ", "-")
	path := storage.GetBookStorage(name)
	file, err := os.Create(path)
	if err != nil {
		logger.Get().Error("Book Send Action: %v", err)
		return "", err
	}
	defer file.Close()

	response, err := config.RequestTelegramBot("getFile", url.Values{"file_id": {document.FileID}})
	if err != nil {
		logger.Get().Error("Making link for downloading file %v", err)
		return "", err
	}

	var body config.ResponseBody
	err = json.NewDecoder(response.Body).Decode(&body)
	if err != nil {
		logger.Get().Error("Reading body getFile %v", err)
		return "", err
	}

	response, err = http.Get(fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", config.Token, body.Result["file_path"].(string)))
	if err != nil {
		logger.Get().Error("Telegram Token getting error: %v", err)
		return "", err
	}
	defer response.Body.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		logger.Get().Error("Telegram Token getting error: %v", err)
		return "", err
	}

	userRepo := model.NewUserService(repository.NewUserRepository())
	user, err := userRepo.Get(model.User{ChatID: b.GetUpdate().FromChat().ID})
	if err != nil {
		return "", err
	}

	repo := model.NewBookService(repository.NewBookRepository())
	book, err := repo.Create(model.Book{
		MessageID:  b.GetUpdate().Message.MessageID,
		Name:       document.FileName,
		Complexity: 0.00,
		Path:       path,
		UserID:     user.ID,
		Status:     model.BOOK_UPLOAD,
	})

	if err != nil {
		return "", err
	}

	err = b.updateStatusUser(b.GetUpdate().FromChat().ID, book.ID)
	if err != nil {
		return "", err
	}

	return textBookSend, nil
}
