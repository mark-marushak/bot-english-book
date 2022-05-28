package action

type StartHandler struct {
	AdaptorTelegramAction
}

//startText = `Hello this bot help you to start read books without any problems!
//			 Let's start with choosing a book you like!
//			 Perhaps you dream about book we don't have, then click upload ahead!
//
//			If you want use this bot,
//			you should share you phone number to active
//			Trial subscribe
//			`
const startText = `Привіт, цей бот допоможе тобі підготувати до читання книги
або просто зрозуміти наскільки ти готовий читати твою книгу.
Якщо ти хочеш скористатись ботом треба зробити ізі реєстрацію
Тобі треба відправити номер телефону і пошту на випадок якщо номер зміниться.
Багато думати не треба просто слідкуй за кнопками внизу та підказками бота.`

func (s StartHandler) Keyboard(i ...interface{}) interface{} {
	return SendPhoneButton
}

func (s StartHandler) Output(i ...interface{}) (string, error) {
	return startText, nil
}
