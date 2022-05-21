package action

type BookUpload struct {
	BaseAction
}

func (b BookUpload) Keyboard(i ...interface{}) interface{} {
	return nil
}

func (b BookUpload) Output(i ...interface{}) string {
	//TODO implement me
	panic("implement me")
}
