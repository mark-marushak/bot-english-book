package account

type TelegramAccount struct {
	id    int64
	name  string
	phone string
}

func (this TelegramAccount) GetName() string {
	return this.name
}
