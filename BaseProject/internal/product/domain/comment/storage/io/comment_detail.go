package io

import (
	"gorm.io/datatypes"
	"time"
)

type CommentDetail struct {
	ID          uint
	ProductID   uint
	UserID      uint
	Description string
	Rating      int
	Name        string
	Avatar      string
	Media       datatypes.JSON
	CreatedAt   time.Time
}
