package websocket

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/real_time/domain/socket"
	chatClient "github.com/eNViDAT0001/Thesis/Backend/internal/real_time/domain/socket/client/chat"
	notifyClient "github.com/eNViDAT0001/Thesis/Backend/internal/real_time/domain/socket/client/notify"
	"github.com/eNViDAT0001/Thesis/Backend/internal/real_time/domain/socket/hub/chat"
	"github.com/eNViDAT0001/Thesis/Backend/internal/real_time/domain/socket/hub/notify"
)

type webSocketHandler struct {
	chatHub      *chat.ChatHub
	chatClient   *chatClient.ChatClient
	notifyHub    *notify.NotifyHub
	notifyClient *notifyClient.NotifyClient
}

func NewWebSocketHandler(chatHub *chat.ChatHub, chatClient *chatClient.ChatClient, notifyHub *notify.NotifyHub, notifyClient *notifyClient.NotifyClient) socket.WebSocketHttpHandler {
	return &webSocketHandler{chatHub: chatHub, chatClient: chatClient, notifyHub: notifyHub, notifyClient: notifyClient}
}
