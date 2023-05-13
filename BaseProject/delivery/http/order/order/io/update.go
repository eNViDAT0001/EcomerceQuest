package io

type UpdateOrderPaymentReq struct {
	OrderIDs   []uint  `json:"order_ids"`
	PaymentID  *string `json:"payment_id"`
	PaymentURL *string `json:"payment_url"`
}
