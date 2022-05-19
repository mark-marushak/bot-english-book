package action

import "github.com/mark-marushak/bot-english-book/pkg/telegram"

type UserStudy struct {
	telegram.Action
}

func (u UserStudy) Send() (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserStudy) Keyboard(i ...interface{}) interface{} {
	//TODO implement me
	panic("implement me")
}

func (u UserStudy) Output(i ...interface{}) string {
	//TODO implement me
	panic("implement me")
}
