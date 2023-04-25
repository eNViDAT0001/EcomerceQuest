package chat

type ChatMessage struct {
}
type ChatHub struct {
	Register   chan *ChatClient
	Unregister chan *ChatClient
	Broadcast  chan *ChatMessage
}

func NewChatHub() *ChatHub {
	return &ChatHub{
		Register:   make(chan *ChatClient),
		Unregister: make(chan *ChatClient),
		Broadcast:  make(chan *ChatMessage, 5),
	}
}
func (h *ChatHub) Run() {

}
