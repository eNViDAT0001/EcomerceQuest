package chat

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/real_time/domain/chat"
)

type chatHandler struct {
	chatUC chat.UseCase
}

func NewChatHandler(chatUC chat.UseCase) chat.HttpHandler {
	return &chatHandler{chatUC: chatUC}
}
