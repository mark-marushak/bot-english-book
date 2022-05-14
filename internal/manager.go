package internal

import (
	"os/user"

	"github.com/mark-marushak/bot-english-book/internal/db"
)

// Manager /*
/*
The Manager must control activity of client
If client don't answer long time (10 minutes), Manager will notification
If client forget about lesson, Manager will send notification to start lesson
If any updates became, Manger will say about them to the client through the bot
*/
type Manager struct{}

func NewManager() *Manager {
	return &Manager{}
}

func (Manager) lookAfterActivity() {
	db.DB().Find(&user.User{})
}

func (Manager) Start() {
}
