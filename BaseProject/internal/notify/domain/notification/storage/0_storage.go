package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/entities"
)

type notificationStorage struct {
}

func (n notificationStorage) CreateNotification(ctx context.Context, input io.NotificationInput) (io.NotificationInput, error) {
	db := wrap_gorm.GetDB()
	err := db.Model(&entities.Notification{}).Create(&input).Error
	return input, err
}

func (n notificationStorage) SeenNotification(ctx context.Context, id uint, userID uint) error {
	db := wrap_gorm.GetDB()
	err := db.Model(&entities.Notification{}).
		Where("id < ?", id).
		Where("user_id = ?", userID).
		Where("seen = ?", false).Update("seen", true).Error
	return err
}

func (n notificationStorage) DeleteByNotificationID(ctx context.Context, id []uint) error {
	db := wrap_gorm.GetDB()
	err := db.Model(&entities.Notification{}).Where("id IN ?", id).Delete(&entities.Notification{}).Error
	return err
}

func (n notificationStorage) ListNotification(ctx context.Context, input io.ListNotifyInput) ([]entities.Notification, error) {
	result := make([]entities.Notification, 0)
	db := wrap_gorm.GetDB()

	query := db.Model(entities.Notification{}).Where("user_id = ?", input.UserID)

	paging_query.SetPagingQuery(&input.Paging, entities.Notification{}.TableName(), query)

	err := query.Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (n notificationStorage) CountListNotification(ctx context.Context, input io.ListNotifyInput) (int64, error) {
	var count int64
	db := wrap_gorm.GetDB()

	query := db.Model(entities.Notification{}).Where("user_id = ?", input.UserID)
	paging_query.SetCountListPagingQuery(&input.Paging, entities.Notification{}.TableName(), query)

	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func NewNotificationStorage() notification.Storage {
	return &notificationStorage{}
}
