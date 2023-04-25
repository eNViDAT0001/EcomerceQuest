package notify

type NotifyMessage struct {
}
type NotifyHub struct {
	Register   chan *NotifyClient
	Unregister chan *NotifyClient
	Broadcast  chan *NotifyMessage
}

func NewNotifyHub() *NotifyHub {
	return &NotifyHub{
		Register:   make(chan *NotifyClient),
		Unregister: make(chan *NotifyClient),
		Broadcast:  make(chan *NotifyMessage, 5),
	}
}
func (h *NotifyHub) Run() {

}
