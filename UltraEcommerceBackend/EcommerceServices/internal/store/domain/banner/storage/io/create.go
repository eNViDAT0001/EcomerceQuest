package io

import "gorm.io/datatypes"

type BannerCreateForm struct {
	ID         uint
	UserID     uint
	Title      string
	Collection string
	Discount   string
	Image      string
	EndTime    datatypes.Date
}
