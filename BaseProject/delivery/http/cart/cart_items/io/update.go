package io

type CartItemUpdateReq struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
