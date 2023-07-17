package io

type OrderReportQuantity struct {
	Date     string `gorm:"date" json:"date"`
	Total    int64  `gorm:"total" json:"total"`
	Quantity int64  `gorm:"quantity" json:"quantity"`
}
