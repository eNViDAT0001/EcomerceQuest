package io

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/order/entities"
	"gorm.io/datatypes"
)

type OrderPreview struct {
	ID       uint                 `json:"id"`
	Name     string               `json:"name"`
	Gender   *bool                `json:"gender"`
	Phone    string               `json:"phone"`
	Province string               `json:"province"`
	District string               `json:"district"`
	Ward     string               `json:"ward"`
	Street   string               `json:"street"`
	Quantity int                  `json:"quantity"`
	Total    int                  `json:"total"`
	Discount int                  `json:"discount"`
	Status   entities.OrderStatus `json:"status"`
	Items    datatypes.JSON       `json:"items"`
}
