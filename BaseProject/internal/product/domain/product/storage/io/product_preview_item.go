package io

import "gorm.io/datatypes"

type ProductPreviewItem struct {
	ID         uint
	ProviderID uint
	CategoryID uint
	UserID     uint
	Name       string
	Price      int64
	Media      datatypes.JSON
	Discount   int
	Rating     float32
}
