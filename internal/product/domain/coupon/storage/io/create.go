package io

type CouponCreateForm struct {
	ID        uint
	ProductID uint
	UserID    uint
	Name      string
	Percent   string
	Fixed     int
}
