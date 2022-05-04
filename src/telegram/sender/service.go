package sender

type SenderService interface {
	Send(id int64, message string)
}
