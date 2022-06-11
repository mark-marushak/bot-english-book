package action

import (
	"fmt"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/internal/repository/gorm"
)

type UserAskEmail struct {
	AdaptorTelegramAction
}

//const userAskEmailText = `Welcome on a board %s,
//			you can either upload one book or choose any book in the bot
//			to start study or just preparation to read full book`
const userAskEmailText = `Welcome on a board %s
Тепер я можу ставитись до тебе як до кліента.
Для початку тобі треба або обрати книгу із моєї біблотеки
Або просто закинь мені книжку у форматі PDF
Також зауважу що файл має бути меншим за 20Мб
Це лише два правила`

func (u UserAskEmail) Keyboard(i ...interface{}) interface{} {
	return AfterRegistrationButton
}

func (u UserAskEmail) Output(i ...interface{}) (string, error) {
	repo := model.NewUserService(gorm.NewUserRepository())
	user, err := repo.Get(model.User{ChatID: u.GetUpdate().FromChat().ID})
	if err != nil {
		return "", fmt.Errorf("UserAskEmail: user wasn't found %v", err)
	}

	user.Email = u.GetUpdate().Message.Text

	err = repo.Update(user)
	if err != nil {
		return "", fmt.Errorf("UserAskEmail: updating %v", err)
	}

	return fmt.Sprintf(userAskEmailText, user.FirstName), nil
}
