package notification

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/entities"
)

type Storage interface {
	Create(ctx context.Context, input io.NotificationCreate) error
	Update(ctx context.Context, id uint, input io.NotificationCreate) error
	Delete(ctx context.Context, input io.NotificationCreate) error
	List(ctx context.Context, input io.ListNotifyInput) ([]entities.Notification, error)
}
