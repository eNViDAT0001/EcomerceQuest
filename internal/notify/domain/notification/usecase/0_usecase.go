package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/entities"
	"gorm.io/gorm"
)

type notificationUseCase struct {
	notifySto notification.Storage
}

func (u *notificationUseCase) SeenNotification(ctx context.Context, id uint, userID uint) error {
	return u.notifySto.SeenNotification(ctx, id, userID)
}
func (u *notificationUseCase) SeenAllNotification(ctx context.Context, id uint, userID uint) error {
	return u.notifySto.SeenAllNotification(ctx, userID)
}

func (u *notificationUseCase) CreateNotification(ctx context.Context, input io.NotificationInput) (io.NotificationInput, error) {
	return u.notifySto.CreateNotification(ctx, input)
}

func (u *notificationUseCase) DeleteByNotificationID(ctx context.Context, id []uint) error {
	return u.notifySto.DeleteByNotificationID(ctx, id)
}

func (u *notificationUseCase) ListNotification(ctx context.Context, input io.ListNotifyInput) ([]entities.Notification, int64, error) {
	total, err := u.notifySto.CountListNotification(ctx, input)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	notifications, err := u.notifySto.ListNotification(ctx, input)
	if err != nil {
		return nil, 0, err
	}

	return notifications, total, err
}

func NewNotificationUseCase(notifySto notification.Storage) notification.UseCase {
	return &notificationUseCase{notifySto: notifySto}
}
