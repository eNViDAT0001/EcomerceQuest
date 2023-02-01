package io

type CreateOrderItemReq struct {
	OrderID         uint
	CategoryID      uint   `json:"category_id"`
	ProductID       uint   `json:"product_id"`
	ProductOptionID uint   `json:"product_option_id"`
	ProviderID      uint   `json:"provider_id"`
	Name            string `json:"name"`
	Price           int    `json:"price"`
	Option          string `json:"option"`
	Quantity        int    `json:"quantity"`
	Discount        int    `json:"discount"`
	Image           string `json:"image"`
}
