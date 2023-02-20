package io

import "gorm.io/datatypes"

type BannerUpdateForm struct {
	Title      *string
	Collection *string
	Discount   *string
	Image      *string
	EndTime    *datatypes.Date
}
