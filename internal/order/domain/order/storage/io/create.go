package io

import "time"

type CreateOrderForm struct {
	ID                uint   `json:"id"`
	UserID            uint   `json:"user_id"`
	ProviderID        uint   `json:"provider_id"`
	Cod               bool   `json:"cod"`
	Name              string `json:"name"`
	Gender            *bool  `json:"gender"`
	Phone             string `json:"phone"`
	Province          string `json:"province"`
	District          string `json:"district"`
	Ward              string `json:"ward"`
	Street            string
	Quantity          int       `json:"quantity"`
	Total             int       `json:"total"`
	Discount          int       `json:"discount"`
	StatusDescription string    `json:"status_description"`
	CreatedAt         time.Time `json:"created_at"`
}
