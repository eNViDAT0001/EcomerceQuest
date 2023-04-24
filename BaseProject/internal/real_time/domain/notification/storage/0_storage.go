package storage

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/real_time/domain/notification"
)

type notificationStorage struct {
}

func NewNotificationStorage() notification.Storage {
	return &notificationStorage{}
}
