package io

type CreateOrderForm struct {
	ID                uint
	UserID            uint
	ProviderID        uint
	Name              string
	Gender            *bool
	Phone             string
	Province          string
	District          string
	Ward              string
	Street            string
	Quantity          int
	Total             int
	Discount          int
	StatusDescription string
}
