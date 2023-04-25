package usecase

import (
	"context"
	"fmt"
	"github.com/eNViDAT0001/Thesis/Backend/external/event_background"
	"github.com/eNViDAT0001/Thesis/Backend/external/html_template"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order/storage/io"
	io2 "github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order_item/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/domain/user"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp"
	io3 "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/usecase/io"
	"gorm.io/gorm"
)

type orderUseCase struct {
	orderSto order.Storage
	userSto  user.Storage
	smtpUC   smtp.UseCase
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
	createdOrders, err := u.orderSto.CreateOrder(ctx, order, items, cartItemsIDs)
	providerIDs := make([]uint, 0)
	for _, createdOrder := range createdOrders {
		providerIDs = append(providerIDs, createdOrder.ProviderID)
	}
	users, err := u.userSto.GetListByProviderID(ctx, providerIDs)
	if err != nil {
		return err
	}

	buyer, err := u.userSto.GetUserDetailByID(ctx, order.UserID)
	if err != nil {
		return err
	}

	jobs := make([]event_background.Job, 0)
	for _, user := range users {
		createdOrder := io.CreateOrderForm{}
		for _, v := range createdOrders {
			if v.UserID == user.ID {
				createdOrder = v
				break
			}
		}

		orderBody := entities.Order{
			SoftDeleteModel: wrap_gorm.SoftDeleteModel{
				CreatedAt: createdOrder.CreatedAt,
				ID:        createdOrder.ID,
			},
			Name:               createdOrder.Name,
			Phone:              createdOrder.Phone,
			Province:           createdOrder.Province,
			District:           createdOrder.District,
			Ward:               createdOrder.Ward,
			Street:             createdOrder.Street,
			Quantity:           createdOrder.Quantity,
			Total:              createdOrder.Total,
			Discount:           createdOrder.Discount,
			Status:             "",
			StatusDescriptions: "",
		}
		email := io3.EmailForm{
			Subject:     fmt.Sprintf("New order from %s", createdOrder.Name),
			Content:     html_template.GetOrderTemplate(orderBody),
			To:          []string{*user.Email},
			Cc:          nil,
			Bcc:         nil,
			AttachFiles: nil,
		}
		jobs = append(jobs, event_background.NewJob(func(ctx context.Context) error {
			return u.smtpUC.SendEmail(ctx, email)
		}))
		email = io3.EmailForm{
			Subject:     "Thanks for your purchase",
			Content:     html_template.GetClientOrderTemplate(orderBody),
			To:          []string{*buyer.Email},
			Cc:          nil,
			Bcc:         nil,
			AttachFiles: nil,
		}
		jobs = append(jobs, event_background.NewJob(func(ctx context.Context) error {
			return u.smtpUC.SendEmail(ctx, email)
		}))
	}
	event_background.GetBackGroundJobs().Group <- event_background.NewGroup(true, jobs...)

	return err
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

func NewOrderUseCase(orderSto order.Storage, userSto user.Storage, smtpUC smtp.UseCase) order.UseCase {
	return &orderUseCase{orderSto: orderSto, userSto: userSto, smtpUC: smtpUC}
}
