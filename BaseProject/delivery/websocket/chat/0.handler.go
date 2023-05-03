package chat

import (
	chat2 "github.com/eNViDAT0001/Thesis/Backend/internal/chat/domain/chat"
)

type socketChatHandler struct {
	chatUC chat2.UseCase
}

func NewSocketChatHandler(chatUC chat2.UseCase) chat2.WebSocketHandler {
	return &socketChatHandler{chatUC: chatUC}
}
