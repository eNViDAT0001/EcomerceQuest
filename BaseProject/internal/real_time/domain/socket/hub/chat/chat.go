package chat

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/real_time/domain/socket/client/chat"
)

type ChatMessage struct {
}
type ChatHub struct {
	Register   chan *chat.ChatClient
	Unregister chan *chat.ChatClient
	Broadcast  chan *ChatMessage
}

func NewChatHub() *ChatHub {
	return &ChatHub{
		Register:   make(chan *chat.ChatClient),
		Unregister: make(chan *chat.ChatClient),
		Broadcast:  make(chan *ChatMessage, 5),
	}
}
func (h *ChatHub) Run() {

}
