package io

import "gorm.io/datatypes"

type BannerDetail struct {
	ID         uint           `json:"id"`
	Title      string         `json:"title"`
	Collection string         `json:"collection"`
	Discount   string         `json:"discount"`
	Image      string         `json:"image"`
	EndTime    datatypes.Date `json:"end_time"`
	Products   datatypes.JSON `json:"products"`
}
