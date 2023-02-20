package order

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/order/domain/order"
)

type orderHandler struct {
	orderUC order.UseCase
}

func NewOrderHandler(orderUC order.UseCase) order.HttpHandler {
	return &orderHandler{orderUC: orderUC}
}
