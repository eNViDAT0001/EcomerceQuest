package io

type BannerCreateForm struct {
	ID        uint
	ProductID uint
	UserID    uint
	Name      string
	Percent   string
	Fixed     int
}
