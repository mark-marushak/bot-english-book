package telegram

/*
ActionService implement three basic functions
Send - resposible for collect keyboard and text for sending
Keyboard - return keyboard based on some data from message
Output - return message text prepared for sending
*/
type ActionRepository interface {
	Keyboard(i ...interface{}) interface{}
	Output(...interface{}) (string, error)
	SetChat(int64)
	GetChat() int64
	SetData(interface{})
	GetData() interface{}
	SetBot(interface{})
	GetBot() interface{}
}

type ActionService interface {
	Keyboard(i ...interface{}) interface{}
	Output(...interface{}) (string, error)
	SetChat(int64)
	GetChat() int64
	SetData(interface{})
	GetData() interface{}
	SetBot(interface{})
	GetBot() interface{}
}

type actionService struct {
	repo ActionRepository
}

func NewAction(repo ActionRepository) ActionService {
	return &actionService{
		repo: repo,
	}
}

func (a actionService) Keyboard(i ...interface{}) interface{} {
	return a.repo.Keyboard(i)
}

func (a actionService) Output(i ...interface{}) (string, error) {
	return a.repo.Output(i)
}

func (a *actionService) SetChat(i int64) {
	a.repo.SetChat(i)
}

func (a actionService) GetChat() int64 {
	return a.repo.GetChat()
}

func (a *actionService) SetData(i interface{}) {
	a.repo.SetData(i)
}

func (a *actionService) GetData() interface{} {
	return a.repo.GetData()
}

func (a *actionService) SetBot(i interface{}) {
	a.repo.SetBot(i)
}

func (a actionService) GetBot() interface{} {
	return a.repo.GetBot()
}
