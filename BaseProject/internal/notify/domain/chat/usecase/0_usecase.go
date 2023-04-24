package usecase

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/chat"
)

type chatUseCase struct {
	chatSto chat.Storage
}

func NewChatUseCase(chatSto chat.Storage) chat.UseCase {
	return &chatUseCase{chatSto: chatSto}
}
