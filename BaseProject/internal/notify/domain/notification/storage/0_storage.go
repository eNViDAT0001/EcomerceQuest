package storage

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification"
)

type notificationStorage struct {
}

func NewNotificationStorage() notification.Storage {
	return &notificationStorage{}
}
