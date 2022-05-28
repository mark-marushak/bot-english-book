package action

type BookUpload struct {
	BaseAction
}

//const textBookUpload = `"Send book please, file must be pdf or txt. For the moment, EnglishBookBot can download files of up to 20MB in size"`
const textBookUpload = `Гаразд ти обрав завантажити свою книгу,
Одразу зауважу файл має бути PDF
Розмір файлу МЕНШЕ 20МБ

Оберай будь-яку книгу із свого арсеналу і давай вже почнємо!`

func (b BookUpload) Keyboard(i ...interface{}) interface{} {
	return DoNothingButton
}

func (b BookUpload) Output(i ...interface{}) (string, error) {
	return textBookUpload, nil
}
