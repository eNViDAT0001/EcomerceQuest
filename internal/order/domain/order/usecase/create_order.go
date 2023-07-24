package usecase

import (
	"context"
	"fmt"
	"github.com/eNViDAT0001/Thesis/Backend/external/html_template"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	notifyIO "github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order/storage/io"
	orderItemsIO "github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order_item/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
	smtpIO "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/socket"
	io2 "github.com/eNViDAT0001/Thesis/Backend/socket/io"
	"strconv"
)

func (u *orderUseCase) CreateOrder(ctx context.Context, order io.CreateOrderForm, items []orderItemsIO.CreateOrderItemForm, cartItemsIDs []uint, couponCodes []string) (createdOrders []io.CreateOrderForm, err error) {
	createdOrders, err = u.orderSto.CreateOrder(ctx, order, items, cartItemsIDs, couponCodes)
	if err != nil {
		return createdOrders, err
	}
	providerIDs := make([]uint, 0)
	for _, createdOrder := range createdOrders {
		providerIDs = append(providerIDs, createdOrder.ProviderID)
	}

	buyer, err := u.userSto.GetUserDetailByID(ctx, order.UserID)
	if err != nil {
		return createdOrders, err
	}

	for _, createdOrder := range createdOrders {
		providerOwner, err := u.userSto.GetDetailByProviderID(ctx, createdOrder.ProviderID)

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
		///
		email := smtpIO.EmailForm{
			Subject:     fmt.Sprintf("New order from %s", buyer.Name),
			Content:     html_template.GetOrderTemplate(orderBody),
			To:          []string{*providerOwner.Email},
			Cc:          nil,
			Bcc:         nil,
			AttachFiles: nil,
		}
		var unSeen = false

		newNotification := notifyIO.NotificationInput{
			ID:      0,
			UserID:  providerOwner.ID,
			Content: "You have a new order from user: " + *buyer.Name,
			Seen:    &unSeen,
			URL:     "/brand-detail/order/" + strconv.Itoa(int(orderBody.ID)),
			Image:   "http://res.cloudinary.com/damzcas3k/image/upload/v1684051785/Product/itl4m7o3jsmtqb2mhtt1.png",
		}

		err = AddNotificationEvent(ctx, newNotification, u.notify.CreateNotification)
		if err != nil {
			return createdOrders, err
		}
		err = u.smtpUC.SendEmail(ctx, email)
		if err != nil {
			return createdOrders, err
		}
		///
		email = smtpIO.EmailForm{
			Subject:     "Thanks for your purchase",
			Content:     html_template.GetClientOrderTemplate(orderBody),
			To:          []string{*buyer.Email},
			Cc:          nil,
			Bcc:         nil,
			AttachFiles: nil,
		}

		newNotification = notifyIO.NotificationInput{
			ID:      0,
			UserID:  buyer.ID,
			Content: "Create order successfully",
			Seen:    &unSeen,
			URL:     "/user/order/" + strconv.Itoa(int(orderBody.ID)),
			Image:   "http://res.cloudinary.com/damzcas3k/image/upload/v1684051785/Product/itl4m7o3jsmtqb2mhtt1.png",
		}
		err = AddNotificationEvent(ctx, newNotification, u.notify.CreateNotification)
		if err != nil {
			return createdOrders, err
		}
		err = u.smtpUC.SendEmail(ctx, email)
		if err != nil {
			return createdOrders, err
		}
	}

	return createdOrders, err
}

func AddNotificationEvent(ctx context.Context, notification notifyIO.NotificationInput, CreateNotification func(context.Context, notifyIO.NotificationInput) (notifyIO.NotificationInput, error)) error {
	newNotification, err := CreateNotification(ctx, notification)
	if err != nil {
		return err
	}
	err = socket.GetManager().EmitNotify(io2.NotificationNew, newNotification, strconv.Itoa(int(newNotification.UserID)))
	if err != nil {
		return err
	}

	return nil
}
