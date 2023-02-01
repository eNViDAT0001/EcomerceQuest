package io

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
	"gorm.io/datatypes"
)

type OrderPreview struct {
	ID       uint
	Name     string
	Gender   *bool
	Phone    string
	Province string
	District string
	Ward     string
	Street   string
	Quantity int
	Total    int
	Discount int
	Status   entities.OrderStatus
	Items    datatypes.JSON
}
