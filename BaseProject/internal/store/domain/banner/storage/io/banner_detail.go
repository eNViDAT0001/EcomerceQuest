package io

import "gorm.io/datatypes"

type BannerDetail struct {
	ID         uint
	Title      string
	Collection string
	Discount   string
	Image      string
	EndTime    datatypes.Date
	Products   datatypes.JSON
}
