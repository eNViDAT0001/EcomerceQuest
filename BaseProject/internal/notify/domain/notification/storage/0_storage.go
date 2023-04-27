package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/entities"
)

type notificationStorage struct {
}

func (n notificationStorage) CreateNotification(ctx context.Context, input io.NotificationInput) error {
	//TODO implement me
	panic("implement me")
}

func (n notificationStorage) UpdateNotification(ctx context.Context, input []io.NotificationInput) error {
	//TODO implement me
	panic("implement me")
}

func (n notificationStorage) DeleteByNotificationID(ctx context.Context, id []uint) error {
	//TODO implement me
	panic("implement me")
}

func (n notificationStorage) ListNotification(ctx context.Context, input io.ListNotifyInput) ([]entities.Notification, error) {
	//TODO implement me
	panic("implement me")
}

func (n notificationStorage) CountListNotification(ctx context.Context, input io.ListNotifyInput) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func NewNotificationStorage() notification.Storage {
	return &notificationStorage{}
}
