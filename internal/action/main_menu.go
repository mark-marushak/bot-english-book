package action

type MainMenu struct {
	AdaptorTelegramAction
}

func (m MainMenu) Keyboard(i ...interface{}) interface{} {
	return MainMenuButton
}

func (m MainMenu) Output(i ...interface{}) (string, error) {
	return "Це головне меню", nil
}
