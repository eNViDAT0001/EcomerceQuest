package notification

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/entities"
)

type UseCase interface {
	CreateNotification(ctx context.Context, input io.NotificationInput) error
	UpdateNotification(ctx context.Context, input []io.NotificationInput) error
	DeleteByNotificationID(ctx context.Context, id []uint) error
	ListNotification(ctx context.Context, input io.ListNotifyInput) ([]entities.Notification, int64, error)
}
