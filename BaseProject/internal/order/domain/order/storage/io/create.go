package io

import "time"

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
	CreatedAt         time.Time
}
