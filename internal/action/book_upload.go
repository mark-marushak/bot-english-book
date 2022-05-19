package action

import "github.com/mark-marushak/bot-english-book/pkg/telegram"

type BookUpload struct {
	telegram.Action
}

func (b BookUpload) Keyboard(i ...interface{}) interface{} {
	return nil
}

func (b BookUpload) Output(i ...interface{}) string {
	//TODO implement me
	panic("implement me")
}
