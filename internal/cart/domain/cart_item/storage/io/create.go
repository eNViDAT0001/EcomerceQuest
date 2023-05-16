package io

type CartItemCreateForm struct {
	CartID    uint
	UserID    uint
	ProductID uint

	ProductOptionID uint
	Quantity        int
}
