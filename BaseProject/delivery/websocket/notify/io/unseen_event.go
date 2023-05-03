package io

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/entities"
)

type NotifyUnseenEvent struct {
	UserID uint              `json:"user_id"`
	Marker int               `json:"marker"`
	Limit  int               `json:"limit"`
	Total  int               `json:"total"`
	Type   paging.PagingType `json:"type"`
	Fields []string          `json:"fields"`
	Search []string          `json:"search"`
	Sort   []string          `json:"sort"`
}
type NewNotifyUnseenEvent struct {
	Notifications []entities.Notification `json:"messages"`
	Paging        paging.Paginator        `json:"paging"`
}
