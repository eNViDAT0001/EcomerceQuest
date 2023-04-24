package websocket

import (
	"github.com/eNViDAT0001/Thesis/Backend/socket"
	"github.com/eNViDAT0001/Thesis/Backend/socket/hub/chat"
	"github.com/eNViDAT0001/Thesis/Backend/socket/hub/notify"
)

type webSocketHandler struct {
	chatHub   *chat.ChatHub
	notifyHub *notify.NotifyHub
}

func NewWebSocketHandler(chatHub *chat.ChatHub, notifyHub *notify.NotifyHub) socket.WebSocketHandler {
	return &webSocketHandler{chatHub: chatHub, notifyHub: notifyHub}
}
