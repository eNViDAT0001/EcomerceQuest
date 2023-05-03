package io

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/entities"
)

type MessageInput struct {
	ID       uint                 `json:"id"`
	UserID   uint                 `json:"user_id"`
	Content  string               `json:"content"`
	ToUserID uint                 `json:"to_user_id"`
	Seen     *bool                `json:"seen"`
	Type     entities.MessageType `json:"type"`
}

type MessageUpdateInput struct {
	UserID   uint
	Content  string
	ToUserID string
	Seen     *bool
	Type     entities.MessageType
}
