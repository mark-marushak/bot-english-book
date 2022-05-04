package telegram

type Bot struct {
	msgsToSend chan response
	done       chan struct {
	}
	Protocol string
	Server   string
}

func (this Bot) Start() {

}
