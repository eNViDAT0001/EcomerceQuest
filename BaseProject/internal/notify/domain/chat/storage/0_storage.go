package storage

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/chat"
)

type chatStorage struct {
}

func NewChatStorage() chat.Storage {
	return &chatStorage{}
}
