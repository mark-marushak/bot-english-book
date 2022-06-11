package action

import (
	"code.sajari.com/docconv"
	"encoding/json"
	"fmt"
	"github.com/mark-marushak/bot-english-book/config"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/internal/repository/gorm"
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

var textBookSend = `Супер, твоя нова книжка додана. 
Почекай повідомлення від мене, 
що книжка готова до навчання
Це зазвичай дуже швидка процедура! 1-2 хвилини`

func (b BookSend) Keyboard(i ...interface{}) interface{} {
	return StartStudyButton
}

func (b BookSend) Output(i ...interface{}) (string, error) {
	var (
		bookRepo = model.NewBookService(gorm.NewBookRepository())
		userRepo = model.NewUserService(gorm.NewUserRepository())
		document = b.GetUpdate().Message.Document
		filepath string
	)

	filepath = document.FileName
	filepath = strings.ReplaceAll(filepath, " ", "-")
	filepath = storage.GetBookStorage(filepath)

	book, err := bookRepo.Get(model.Book{
		Name: document.FileName,
		Path: filepath,
	})

	if err != nil {
		return "", err
	}

	if book.ID > 0 {
		return "Ця книжка вже є в нащі бібліотеці", nil
	}

	filepath, err = b.getFile(filepath)
	if err != nil {
		if os.Remove(filepath) != nil {
			logger.Get().Error("File removing error: %v", err)
		}
		return "Це є не допустимий формат. Будь ласка перевірте правельність формату який відправляєте. Формат має бути PDF", nil
	}

	user, err := userRepo.Get(model.User{ChatID: b.GetUpdate().FromChat().ID})
	if err != nil {
		return "", err
	}

	book, err = bookRepo.Create(model.Book{
		MessageID:  b.GetUpdate().Message.MessageID,
		Name:       document.FileName,
		Complexity: 0.00,
		Path:       filepath,
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

func (b BookSend) getFile(filepath string) (string, error) {
	var document = b.GetUpdate().Message.Document
	var err error

	file, err := os.Create(filepath)
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
	defer response.Body.Close()

	var body config.ResponseBody
	err = json.NewDecoder(response.Body).Decode(&body)
	if err != nil {
		logger.Get().Error("Reading body getFile %v", err)
		return "", err
	}

	err = b.DownloadFile(body.Result["file_path"].(string), file)
	if err != nil {
		return "", err
	}

	text, err := docconv.ConvertPath(filepath)
	if err != nil {
		return "", err
	}
	if len(text.Body) > 0 {
		return filepath, nil
	}

	return "", fmt.Errorf("file is wrong")
}

func (b BookSend) DownloadFile(remoteFilePath string, file *os.File) error {
	response, err := http.Get(fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", config.Token, remoteFilePath))
	if err != nil {
		logger.Get().Error("Telegram Token getting error: %v", err)
		return err
	}
	defer response.Body.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		logger.Get().Error("Telegram Token getting error: %v", err)
		return err
	}

	return nil
}
