package notification

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/entities"
)

type UseCase interface {
	CreateNotification(ctx context.Context, input io.NotificationInput) (io.NotificationInput, error)
	SeenNotification(ctx context.Context, id uint, userID uint) error
	SeenAllNotification(ctx context.Context, userID uint) error
	DeleteByNotificationID(ctx context.Context, id []uint) error
	ListNotification(ctx context.Context, input io.ListNotifyInput) ([]entities.Notification, int64, error)
}
