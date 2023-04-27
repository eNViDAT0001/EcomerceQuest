package usecase

import (
	"context"
	notification2 "github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/entities"
)

type notificationUseCase struct {
	notifySto notification2.Storage
}

func (n notificationUseCase) CreateNotification(ctx context.Context, input io.NotificationInput) error {
	//TODO implement me
	panic("implement me")
}

func (n notificationUseCase) UpdateNotification(ctx context.Context, input []io.NotificationInput) error {
	//TODO implement me
	panic("implement me")
}

func (n notificationUseCase) DeleteByNotificationID(ctx context.Context, id []uint) error {
	//TODO implement me
	panic("implement me")
}

func (n notificationUseCase) ListNotification(ctx context.Context, input io.ListNotifyInput) ([]entities.Notification, int64, error) {
	//TODO implement me
	panic("implement me")
}

func NewNotificationUseCase(notifySto notification2.Storage) notification2.UseCase {
	return &notificationUseCase{notifySto: notifySto}
}
