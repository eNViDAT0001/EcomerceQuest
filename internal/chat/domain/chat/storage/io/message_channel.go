package io

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/entities"
)

type MessageChannel struct {
	ChatRoomID uint                 `gorm:"column:chat_room_id" json:"chat_room_id"`
	Name       string               `gorm:"column:name" json:"name"`
	Avatar     string               `gorm:"column:avatar" json:"avatar"`
	UserID     uint                 `gorm:"column:user_id" json:"user_id"`
	Content    string               `gorm:"column:content" json:"content"`
	ToUserID   uint                 `gorm:"column:to_user_id" json:"to_user_id"`
	Seen       *bool                `gorm:"column:seen" json:"seen"`
	Type       entities.MessageType `gorm:"column:type" json:"type"`
}
type ChatRoom struct {
	ID         uint                 `gorm:"column:id" json:"id"`
	FromUserID uint                 `gorm:"column:from_user_id" json:"from_user_id"`
	ToUserID   uint                 `gorm:"column:to_user_id" json:"to_user_id"`
	Content    string               `gorm:"column:content" json:"content"`
	Seen       bool                 `gorm:"column:seen" json:"seen"`
	Type       entities.MessageType `gorm:"column:type" json:"type"`
	CreatedAt  string               `gorm:"column:created_at" json:"created_at"`

	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}
type Channel struct {
	UserID   uint `gorm:"column:user_id" json:"user_id"`
	ToUserID uint `gorm:"column:to_user_id" json:"to_user_id"`
}
