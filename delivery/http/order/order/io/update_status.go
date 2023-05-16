package io

import "github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"

type UpdateOrderStatusReq struct {
	Status entities.OrderStatus `json:"status"`
	Image  string               `json:"image"`
}
