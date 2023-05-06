package io

type BannerCreateForm struct {
	ID         uint
	UserID     uint
	Title      string
	Collection string
	Discount   int
	Image      string
	EndTime    string
}
