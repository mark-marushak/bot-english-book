package chain_of_command

type Handler interface {
	SetNext(h Handler)
	Handle(request string)
}

type BaseHandler struct {
	next Handler
}

func (this BaseHandler) SetNext(h Handler) {
	this.next = h
}

func (this BaseHandler) Handle(request string) {
	if this.next != nil {
		this.next.Handle(request)
	}
}
