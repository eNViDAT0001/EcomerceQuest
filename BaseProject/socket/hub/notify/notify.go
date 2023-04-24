package notify

import (
	"github.com/eNViDAT0001/Thesis/Backend/socket/client/notify"
)

type NotifyMessage struct {
}
type NotifyHub struct {
	Register   chan *notify.NotifyClient
	Unregister chan *notify.NotifyClient
	Broadcast  chan *NotifyMessage
}

func NewNotifyHub() *NotifyHub {
	return &NotifyHub{
		Register:   make(chan *notify.NotifyClient),
		Unregister: make(chan *notify.NotifyClient),
		Broadcast:  make(chan *NotifyMessage, 5),
	}
}
func (h *NotifyHub) Run() {

}
