package io

type CouponCreateForm struct {
	ID      uint
	UserID  uint
	Code    string
	Percent int
	Fixed   float64
}
type CouponDetailCreateForm struct {
	ID        uint
	CouponID  uint
	ProductID uint
	Total     int
}
