package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order/storage/io"
	io2 "github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order_item/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
	"gorm.io/gorm"
)

type orderUseCase struct {
	orderSto order.Storage
}

func (u *orderUseCase) ListByUserID(ctx context.Context, userID uint, input paging.ParamsInput) (orders []entities.Order, total int64, err error) {
	total, err = u.orderSto.CountListByUserID(ctx, userID, input)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	orders, err = u.orderSto.ListByUserID(ctx, userID, input)
	if err != nil {
		return nil, 0, err
	}

	return orders, total, err
}

func (u *orderUseCase) ListPreviewByUserID(ctx context.Context, userID uint, input paging.ParamsInput) (orders []io.OrderPreview, total int64, err error) {
	total, err = u.orderSto.CountPreviewByUserID(ctx, userID, input)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	orders, err = u.orderSto.ListPreviewByUserID(ctx, userID, input)
	if err != nil {
		return nil, 0, err
	}

	return orders, total, err
}

func (u *orderUseCase) ListByProviderID(ctx context.Context, providerID uint, input paging.ParamsInput) (orders []entities.Order, total int64, err error) {
	total, err = u.orderSto.CountByProviderID(ctx, providerID, input)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	orders, err = u.orderSto.ListByProviderID(ctx, providerID, input)
	if err != nil {
		return nil, 0, err
	}

	return orders, total, err
}

func (u *orderUseCase) ListPreviewByProviderID(ctx context.Context, providerID uint, input paging.ParamsInput) (orders []io.OrderPreview, total int64, err error) {
	total, err = u.orderSto.CountPreviewByProviderID(ctx, providerID, input)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	orders, err = u.orderSto.ListPreviewByProviderID(ctx, providerID, input)
	if err != nil {
		return nil, 0, err
	}

	return orders, total, err
}

func (u *orderUseCase) GetByOrderID(ctx context.Context, orderID uint) (entities.Order, error) {
	return u.orderSto.GetByOrderID(ctx, orderID)
}

func (u *orderUseCase) CreateOrder(ctx context.Context, order io.CreateOrderForm, items []io2.CreateOrderItemForm, cartItemsIDs []uint) (err error) {
	return u.orderSto.CreateOrder(ctx, order, items, cartItemsIDs)
}

func (u *orderUseCase) UpdateOrderStatus(ctx context.Context, orderID uint, status entities.OrderStatus) error {
	return u.orderSto.UpdateOrderStatus(ctx, orderID, status)
}

func (u *orderUseCase) CancelOrder(ctx context.Context, orderID uint) error {
	return u.orderSto.CancelOrder(ctx, orderID)
}

func (u *orderUseCase) DeleteOrder(ctx context.Context, orderID uint) error {
	return u.orderSto.DeleteOrder(ctx, orderID)
}

func NewOrderUseCase(orderSto order.Storage) order.UseCase {
	return &orderUseCase{orderSto: orderSto}
}
