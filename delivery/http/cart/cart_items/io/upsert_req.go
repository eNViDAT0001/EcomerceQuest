package io

type CartItemCreateReq struct {
	UserID    uint
	ProductID uint

	ProductOptionID uint `json:"product_option_id"`
	Quantity        int  `json:"quantity"`
}
