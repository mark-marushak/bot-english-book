package action

type BookUpload struct {
	BaseAction
}

func (b BookUpload) Keyboard(i ...interface{}) interface{} {
	return nil
}

func (b BookUpload) Output(i ...interface{}) (string, error) {
	return "Send book please, file must be pdf or txt. For the moment, EnglishBookBot can download files of up to 20MB in size", nil
}
