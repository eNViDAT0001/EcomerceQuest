package io

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/entities"
)

type MessageChannel struct {
	ChannelID uint                 `gorm:"column:channel_id" json:"channel_id"`
	Name      string               `gorm:"column:name" json:"name"`
	Avatar    string               `gorm:"column:avatar" json:"avatar"`
	UserID    uint                 `gorm:"column:user_id" json:"user_id"`
	Content   string               `gorm:"column:content" json:"content"`
	ToUserID  string               `gorm:"column:to_user_id" json:"to_user_id"`
	Seen      *bool                `gorm:"column:seen" json:"seen"`
	Type      entities.MessageType `gorm:"column:type" json:"type"`
}
type Channel struct {
	UserID   uint `gorm:"column:user_id" json:"user_id"`
	ToUserID uint `gorm:"column:to_user_id" json:"to_user_id"`
}
