package io

import "gorm.io/datatypes"

type CouponDetail struct {
	ID        uint           `json:"id"`
	ProductID uint           `json:"product_id"`
	UserID    uint           `json:"user_id"`
	Name      string         `json:"name"`
	Percent   string         `json:"percent"`
	Fixed     int            `json:"fixed"`
	Products  datatypes.JSON `json:"products"`
}
