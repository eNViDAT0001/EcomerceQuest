package io

type CouponCreateForm struct {
	ID      uint
	UserID  uint
	Code    string
	Percent int
	Fixed   float64
}
type CouponDetailCreateForm struct {
	CouponID  uint
	ProductID uint
	Total     int
}
