package io

import "github.com/eNViDAT0001/Thesis/Backend/internal/notify/entities"

type MessageCreateReq struct {
	ID       uint
	UserID   uint                 `json:"user_id"`
	Content  string               `json:"content"`
	ToUserID string               `json:"to_user_id"`
	Seen     *bool                `json:"seen"`
	Type     entities.MessageType `json:"type"`
}

type MessageUpdateReq struct {
	UserID   uint                 `json:"user_id"`
	Content  string               `json:"content"`
	ToUserID string               `json:"to_user_id"`
	Seen     *bool                `json:"seen"`
	Type     entities.MessageType `json:"type"`
}
