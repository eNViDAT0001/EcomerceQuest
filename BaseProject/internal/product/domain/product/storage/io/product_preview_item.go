package io

import "gorm.io/datatypes"

type ProductPreviewItem struct {
	ID         uint           `json:"id"`
	ProviderID uint           `json:"provider_id"`
	CategoryID uint           `json:"category_id"`
	UserID     uint           `json:"user_id"`
	Name       string         `json:"name"`
	Price      int64          `json:"price"`
	Media      datatypes.JSON `json:"media"`
	Discount   int            `json:"discount"`
	Rating     float32        `json:"rating"`
}
