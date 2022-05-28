package action

type BaseAction struct {
	ChatID int64
	Data   interface{}
	Bot    interface{}
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

func (u *BaseAction) SetBot(i interface{}) {
	u.Bot = i
}

func (u BaseAction) GetBot() interface{} {
	return u.Bot
}
