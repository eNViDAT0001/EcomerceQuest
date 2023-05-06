package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eNViDAT0001/Thesis/Backend/external/event_background"
	"github.com/eNViDAT0001/Thesis/Backend/external/html_template"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	notifyIO "github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order/storage/io"
	orderItemsIO "github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order_item/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
	smtpIO "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/usecase/io"
	"github.com/eNViDAT0001/Thesis/Backend/socket"
	socketIO "github.com/eNViDAT0001/Thesis/Backend/socket/io"
	"strconv"
)

func (u *orderUseCase) CreateOrder(ctx context.Context, order io.CreateOrderForm, items []orderItemsIO.CreateOrderItemForm, cartItemsIDs []uint) (err error) {
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
		///
		email := smtpIO.EmailForm{
			Subject:     fmt.Sprintf("New order from %s", createdOrder.Name),
			Content:     html_template.GetOrderTemplate(orderBody),
			To:          []string{*user.Email},
			Cc:          nil,
			Bcc:         nil,
			AttachFiles: nil,
		}
		var unSeen = false

		jobs = append(jobs, event_background.NewJob(func(ctx context.Context) error {
			newNotification := notifyIO.NotificationInput{
				ID:      0,
				UserID:  user.ID,
				Content: "You have a new order from user: " + *user.Name,
				Seen:    &unSeen,
				URL:     "api/v1/orders/" + strconv.Itoa(int(orderBody.ID)),
			}

			err = AddNotificationEvent(ctx, newNotification, u.notify.CreateNotification)
			if err != nil {
				return err
			}
			return u.smtpUC.SendEmail(ctx, email)
		}))
		///
		email = smtpIO.EmailForm{
			Subject:     "Thanks for your purchase",
			Content:     html_template.GetClientOrderTemplate(orderBody),
			To:          []string{*buyer.Email},
			Cc:          nil,
			Bcc:         nil,
			AttachFiles: nil,
		}

		jobs = append(jobs, event_background.NewJob(func(ctx context.Context) error {
			newNotification := notifyIO.NotificationInput{
				ID:      0,
				UserID:  buyer.ID,
				Content: "Create order successfully",
				Seen:    &unSeen,
				URL:     "api/v1/orders/" + strconv.Itoa(int(orderBody.ID)),
			}
			err = AddNotificationEvent(ctx, newNotification, u.notify.CreateNotification)
			if err != nil {
				return err
			}

			return u.smtpUC.SendEmail(ctx, email)
		}))
	}
	event_background.GetBackGroundJobs().Group <- event_background.NewGroup(true, jobs...)

	return err
}

func AddNotificationEvent(ctx context.Context, notification notifyIO.NotificationInput, CreateNotification func(context.Context, notifyIO.NotificationInput) (notifyIO.NotificationInput, error)) error {
	newNotification, err := CreateNotification(ctx, notification)
	if err != nil {
		return err
	}
	data, err := json.Marshal(newNotification)
	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}

	socketManager := socket.GetManager()
	socketManager.Lock()
	if _, ok := socketManager.Clients[strconv.Itoa(int(newNotification.UserID))]; ok {
		socketManager.Clients[strconv.Itoa(int(newNotification.UserID))].AddEvent(socketIO.Event{
			Type:    socketIO.NotificationNew,
			Payload: data,
		})
	}
	socketManager.Unlock()
	return nil
}
