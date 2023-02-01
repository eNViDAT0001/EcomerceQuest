package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order_item"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
)

type orderItemUseCase struct {
	orderItemSto order_item.Storage
}

func (u *orderItemUseCase) ListByOrderID(ctx context.Context, orderID uint) ([]entities.OrderItem, error) {
	return u.orderItemSto.ListByOrderID(ctx, orderID)
}

func NewOrderItemUseCase(orderItemSto order_item.Storage) order_item.UseCase {
	return &orderItemUseCase{orderItemSto: orderItemSto}
}
