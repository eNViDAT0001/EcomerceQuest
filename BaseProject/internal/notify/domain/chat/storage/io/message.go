package io

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/entities"
)

type MessageInput struct {
	ID       uint
	UserID   uint
	Content  string
	ToUserID string
	Seen     *bool
	Type     entities.MessageType
}

type MessageUpdateInput struct {
	UserID   uint
	Content  string
	ToUserID string
	Seen     *bool
	Type     entities.MessageType
}
