package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order_item"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/payment"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/domain/user"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp"
	"gorm.io/gorm"
)

type orderUseCase struct {
	orderSto   order.Storage
	userSto    user.Storage
	smtpUC     smtp.UseCase
	paymentUC  payment.UseCase
	orderItems order_item.Storage
	notify     notification.UseCase
}

func (u *orderUseCase) GetFullDetailByOrderID(ctx context.Context, orderID uint) (io.OrderFullDetail, error) {

	ordered, err := u.orderSto.GetByOrderID(ctx, orderID)
	if err != nil {
		return io.OrderFullDetail{}, err
	}

	orderItems, err := u.orderItems.ListByOrderID(ctx, orderID)
	if err != nil {
		return io.OrderFullDetail{}, err
	}

	result := io.OrderFullDetail{
		ID:        orderID,
		Name:      ordered.Name,
		Gender:    ordered.Gender,
		Phone:     ordered.Phone,
		Province:  ordered.Province,
		District:  ordered.District,
		Ward:      ordered.Ward,
		Street:    ordered.Street,
		Quantity:  ordered.Quantity,
		Total:     ordered.Total,
		Discount:  ordered.Discount,
		Status:    ordered.Status,
		Items:     orderItems,
		CreatedAt: ordered.CreatedAt,
	}

	if ordered.PaymentID != nil {
		paymentInfo, err := u.paymentUC.GetPaymentByID(ctx, *ordered.PaymentID)
		if err != nil {
			return io.OrderFullDetail{}, err
		}
		result.Payment = &paymentInfo
	}
	return result, nil
}

func (u *orderUseCase) DeleteOrders(ctx context.Context, ids []uint) error {
	return u.orderSto.DeleteOrders(ctx, ids)
}

func (u *orderUseCase) ListInvalidOrder(ctx context.Context) (unPayOrders []entities.Order, unConfirmedOrders []entities.Order, err error) {
	unPayOrders, err = u.orderSto.ListUnPayOrder(ctx)
	if err != nil {
		return nil, nil, err
	}
	unConfirmedOrders, err = u.orderSto.ListUnConfirmedDeliveredOrder(ctx)
	if err != nil {
		return nil, nil, err
	}

	return unPayOrders, unConfirmedOrders, nil
}

func (u *orderUseCase) UpdateOrder(ctx context.Context, orderID uint, input io.UpdateOrderForm) error {
	return u.orderSto.UpdateOrder(ctx, orderID, input)
}

func (u *orderUseCase) UpdateOrderPayment(ctx context.Context, orderID []uint, input io.UpdateOrderPaymentForm) error {
	return u.orderSto.UpdateOrderPayment(ctx, orderID, input)
}

func (u *orderUseCase) VerifyDeliveredOrder(ctx context.Context, orderID uint, userID uint) error {
	return u.orderSto.VerifyDeliveredOrder(ctx, orderID, userID)
}

func (u *orderUseCase) UpdateDeliveredOrderStatus(ctx context.Context, orderID uint, image string) error {
	return u.orderSto.UpdateDeliveredOrderStatus(ctx, orderID, image)
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

func (u *orderUseCase) ListPreview(ctx context.Context, input paging.ParamsInput) (orders []io.OrderPreview, total int64, err error) {
	total, err = u.orderSto.CountPreview(ctx, input)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	orders, err = u.orderSto.ListPreview(ctx, input)
	if err != nil {
		return nil, 0, err
	}

	return orders, total, err
}
func (u *orderUseCase) List(ctx context.Context, input paging.ParamsInput) (orders []entities.Order, total int64, err error) {
	total, err = u.orderSto.CountList(ctx, input)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	orders, err = u.orderSto.List(ctx, input)
	if err != nil {
		return nil, 0, err
	}

	return orders, total, err
}

func (u *orderUseCase) GetByOrderID(ctx context.Context, orderID uint) (entities.Order, error) {
	return u.orderSto.GetByOrderID(ctx, orderID)
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

func NewOrderUseCase(
	orderSto order.Storage,
	userSto user.Storage,
	smtpUC smtp.UseCase,
	paymentUC payment.UseCase,
	orderItems order_item.Storage,
	notify notification.UseCase) order.UseCase {
	return &orderUseCase{
		orderSto:   orderSto,
		userSto:    userSto,
		smtpUC:     smtpUC,
		paymentUC:  paymentUC,
		orderItems: orderItems,
		notify:     notify}
}
