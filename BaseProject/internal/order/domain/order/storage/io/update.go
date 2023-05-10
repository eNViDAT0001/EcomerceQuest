package io

import (
	"gorm.io/datatypes"
)

type UpdateOrderForm struct {
	UserID             uint            `json:"user_id"`
	ProviderID         uint            `json:"provider_id"`
	PaymentID          *string         `json:"payment_id"`
	PaymentURL         *string         `json:"payment_url"`
	Name               string          `json:"name"`
	Gender             *bool           `json:"gender"`
	Phone              string          `json:"phone"`
	Province           string          `json:"province"`
	District           string          `json:"district"`
	Ward               string          `json:"ward"`
	Street             string          `json:"street"`
	Quantity           int             `json:"quantity"`
	Total              int             `json:"total"`
	Discount           int             `json:"discount"`
	StatusDescriptions string          `json:"status_descriptions"`
	DeliveredImage     *string         `json:"delivered_image"`
	VerifyDelivered    *bool           `json:"verify_delivered"`
	DeliveredDate      *datatypes.Time `json:"delivered_date"`
}
