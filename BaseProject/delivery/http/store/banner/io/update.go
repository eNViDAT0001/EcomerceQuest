package io

import "gorm.io/datatypes"

type BannerUpdateReq struct {
	Title         *string         `json:"title"`
	Collection    *string         `json:"collection"`
	Discount      *string         `json:"discount"`
	Image         *string         `json:"image"`
	EndTime       *datatypes.Date `json:"end_time"`
	ProductIDsIN  []uint          `json:"product_ids_in"`
	ProductIDsOUT []uint          `json:"product_ids_out"`
}
