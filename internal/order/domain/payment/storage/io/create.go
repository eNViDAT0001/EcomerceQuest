package io

type CreatePaymentForm struct {
	ID        string `json:"id"`
	AccountID string `json:"account_id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Status    *bool  `json:"status"`
}
type UpdatePaymentForm struct {
	AccountID uint   `json:"account_id"`
	Email     uint   `json:"email"`
	Name      string `json:"name"`
	Status    *bool  `json:"status"`
}
