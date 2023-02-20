package io

type CreateOrderReq struct {
	ID                uint
	UserID            uint                 `json:"user_id"`
	ProviderID        uint                 `json:"provider_id"`
	Name              string               `json:"name"`
	Gender            *bool                `json:"gender"`
	Phone             string               `json:"phone"`
	Province          string               `json:"province"`
	District          string               `json:"district"`
	Ward              string               `json:"ward"`
	Street            string               `json:"street"`
	Quantity          int                  `json:"quantity"`
	Total             int                  `json:"total"`
	Discount          int                  `json:"discount"`
	StatusDescription string               `json:"status_description"`
	Items             []CreateOrderItemReq `json:"items"`
	CartItemsIDS      []uint               `json:"cart_items_ids"`
}
