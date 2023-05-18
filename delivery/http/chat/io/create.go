package io

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/entities"
)

type MessageCreateReq struct {
	ID         uint
	ChatRoomID uint                 `json:"chat_room_id"`
	Content    string               `json:"content"`
	ToUserID   uint                 `json:"to_user_id"`
	FromUserID uint                 `json:"from_user_id"`
	Seen       *bool                `json:"seen"`
	Type       entities.MessageType `json:"type"`
}

type MessageUpdateReq struct {
	UserID   uint                 `json:"user_id"`
	Content  string               `json:"content"`
	ToUserID uint                 `json:"to_user_id"`
	Seen     *bool                `json:"seen"`
	Type     entities.MessageType `json:"type"`
}
