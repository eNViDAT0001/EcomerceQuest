package order

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order/storage/io"
	io2 "github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order_item/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
)

type Storage interface {
	ListByUserID(ctx context.Context, userID uint, input paging.ParamsInput) ([]entities.Order, error)
	CountListByUserID(ctx context.Context, userID uint, input paging.ParamsInput) (total int64, err error)
	ListPreviewByUserID(ctx context.Context, userID uint, input paging.ParamsInput) ([]io.OrderPreview, error)
	CountPreviewByUserID(ctx context.Context, userID uint, input paging.ParamsInput) (total int64, err error)

	ListByProviderID(ctx context.Context, providerID uint, input paging.ParamsInput) ([]entities.Order, error)
	CountByProviderID(ctx context.Context, providerID uint, input paging.ParamsInput) (total int64, err error)

	ListPreviewByProviderID(ctx context.Context, providerID uint, input paging.ParamsInput) ([]io.OrderPreview, error)
	CountPreviewByProviderID(ctx context.Context, providerID uint, input paging.ParamsInput) (total int64, err error)

	GetByOrderID(ctx context.Context, orderID uint) (entities.Order, error)

	CreateOrder(ctx context.Context, order io.CreateOrderForm, items []io2.CreateOrderItemForm, cartItemsIDs []uint) (createdOrders []io.CreateOrderForm, err error)

	UpdateOrderStatus(ctx context.Context, orderID uint, status entities.OrderStatus) error
	CancelOrder(ctx context.Context, orderID uint) error
	DeleteOrder(ctx context.Context, orderID uint) error
}
