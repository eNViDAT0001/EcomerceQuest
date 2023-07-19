package io

type CouponDetail struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	UserID    uint   `json:"user_id"`
	Name      string `json:"name"`
	Percent   string `json:"percent"`
	Fixed     int    `json:"fixed"`
}
