package usecase

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/real_time/domain/chat"
)

type chatUseCase struct {
	chatSto chat.Storage
}

func NewChatUseCase(chatSto chat.Storage) chat.Storage {
	return &chatUseCase{chatSto: chatSto}
}
