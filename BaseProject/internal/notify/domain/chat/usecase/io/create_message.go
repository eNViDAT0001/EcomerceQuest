package io

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/entities"
)

type Message struct {
	ID       uint
	UserID   uint
	Content  string
	ToUserID string
	Seen     bool
	Type     entities.MessageType
}
