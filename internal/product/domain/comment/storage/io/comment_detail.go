package io

import (
	"gorm.io/datatypes"
	"time"
)

type CommentDetail struct {
	ID          uint           `json:"id"`
	ProductID   uint           `json:"product_id"`
	UserID      uint           `json:"user_id"`
	Description string         `json:"description"`
	Rating      int            `json:"rating"`
	Name        string         `json:"name"`
	Avatar      string         `json:"avatar"`
	Media       datatypes.JSON `json:"media"`
	CreatedAt   time.Time      `json:"created_at"`
}
