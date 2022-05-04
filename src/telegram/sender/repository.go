package sender

type SenderRepository interface {
	Send(id int64, message string)
}
