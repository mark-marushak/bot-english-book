package sender

type SendService struct {
	sender     SenderService
	protoParam interface{}
}

func NewSender(repository SenderRepository) *SendService {
	return &SendService{
		sender: repository,
	}
}

func (s SendService) Send(id int64, message string) {
	s.sender.Send(id, message)
}
