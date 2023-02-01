package io

type CreateOrderItemForm struct {
	OrderID         uint
	ProductID       uint
	ProductOptionID uint
	ProviderID      uint
	Name            string
	Price           int
	Option          string
	Quantity        int
	Discount        int
	Image           string
}
