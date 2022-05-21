package action

type BaseAction struct {
	ChatID int64
	Data   interface{}
}

func (u *BaseAction) SetChat(i int64) {
	u.ChatID = i
}

func (u BaseAction) GetChat() int64 {
	return u.ChatID
}

func (u *BaseAction) SetData(i interface{}) {
	u.Data = i
}

func (u BaseAction) GetData() interface{} {
	return u.Data
}
