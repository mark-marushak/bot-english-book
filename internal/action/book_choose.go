package action

import (
	"github.com/mark-marushak/bot-english-book/logger"
	"regexp"
	"strconv"
)

type BookChoose struct {
	AdaptorTelegramAction
}

var textBookChoose = `Дуже гарний вибір,
ця книжка вже підготовлена для підготовки тебе до неї :)
Тисни на кнопку і почнемо
`

func (b BookChoose) Keyboard(i ...interface{}) interface{} {
	return StartStudyButton
}

func (b BookChoose) Output(i ...interface{}) (string, error) {
	bookID := b.GetUpdate().CallbackData()
	result := string(regexp.MustCompile("\\d+").Find([]byte(bookID)))

	id, err := strconv.Atoi(result)
	if err != nil {
		logger.Get().Error("BookChoose: parse id error: %v", err)
		return "", nil
	}

	err = b.updateStatusUser(b.GetUpdate().FromChat().ID, uint(id))
	if err != nil {
		return "", err
	}

	return textBookChoose, nil
}
