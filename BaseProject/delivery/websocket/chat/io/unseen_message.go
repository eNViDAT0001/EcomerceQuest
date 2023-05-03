package io

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/entities"
)

type ChatConnectEvent struct {
	UserID   uint `json:"user_id"`
	ToUserID uint `json:"to_user_id"`

	Marker int               `json:"marker"`
	Limit  int               `json:"limit"`
	Total  int               `json:"total"`
	Type   paging.PagingType `json:"type"`
	Fields []string          `json:"fields"`
	Search []string          `json:"search"`
	Sort   []string          `json:"sort"`
}
type NewChatConnectEvent struct {
	Messages []entities.Message `json:"messages"`
	Paging   paging.Paginator   `json:"paging"`
}
