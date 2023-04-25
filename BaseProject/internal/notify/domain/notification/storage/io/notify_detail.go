package io

import "github.com/eNViDAT0001/Thesis/Backend/internal/notify/entities"

type NotificationDetail struct {
	ID      uint                      `json:"id"`
	UserID  uint                      `json:"user_id"`
	Content string                    `json:"content"`
	Seen    bool                      `json:"seen"`
	Type    entities.NotificationType `json:"type"`
	URL     string                    `json:"url"`
}
type NotificationCreate struct {
	ID      uint                      `json:"id"`
	UserID  uint                      `json:"user_id"`
	Content string                    `json:"content"`
	Seen    bool                      `json:"seen"`
	Type    entities.NotificationType `json:"type"`
	URL     string                    `json:"url"`
}
