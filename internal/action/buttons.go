package action

import tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const StartStudy = "Навчатись"
const NextLesson = "Наступне завдання"
const OpenLibrary = "Відкрити бібліотеку"
const UploadBook = "Завантажити книгу"
const StartRegistration = "Почати реєстрацію"
const BackToMainMenu = "Головне меню"
const DoNothing = "Не знаю що робити"

var MainMenuButton = tg.NewReplyKeyboard(
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(StartStudy),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(OpenLibrary),
		tg.NewKeyboardButton(UploadBook),
	),
)

var SendPhoneButton = tg.NewOneTimeReplyKeyboard(
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButtonContact(StartRegistration),
	),
)

var BackToMainMenuButton = tg.NewReplyKeyboard(
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(BackToMainMenu),
	),
)

var AfterRegistrationButton = tg.NewOneTimeReplyKeyboard(
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(UploadBook),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(OpenLibrary),
	),
)

var DoNothingButton = tg.NewReplyKeyboard(
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(DoNothing),
	),
)

var NextLessonButton = tg.NewOneTimeReplyKeyboard(
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(NextLesson),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(BackToMainMenu),
	),
)

var StartStudyButton = tg.NewOneTimeReplyKeyboard(
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(StartStudy),
	),
)
