package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order_item"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
	"time"
)

type orderStorage struct {
	orderItemSto order_item.Storage
}

func (s orderStorage) ListUnPayOrder(ctx context.Context) ([]entities.Order, error) {
	result := make([]entities.Order, 0)
	db := wrap_gorm.GetDB()
	err := db.Model(entities.Order{}).Where("payment_id IS NULL").Find(&result).Error

	return result, err
}

func (s orderStorage) ListUnConfirmedDeliveredOrder(ctx context.Context) ([]entities.Order, error) {
	result := make([]entities.Order, 0)
	db := wrap_gorm.GetDB()
	deadline := time.Now().Add(-24 * time.Hour)
	err := db.Model(entities.Order{}).
		Where("status = ?", entities.DeliveredOrder).
		Where("verify_delivered = 0").
		Where("delivered_date < ?", deadline).
		Find(&result).Error

	return result, err
}

func (s orderStorage) VerifyDeliveredOrder(ctx context.Context, orderID uint, userID uint) error {
	db := wrap_gorm.GetDB()

	err := db.Model(entities.Order{}).
		Where("id = ? AND user_id = ?", orderID, userID).
		Update("verify_delivered", 1).
		Error

	return err
}

type UpdateOrderStatus struct {
	Status        entities.OrderStatus `json:"status"`
	Image         string               `json:"image"`
	DeliveredDate time.Time            `json:"delivered_date"`
}

func (s orderStorage) UpdateDeliveredOrderStatus(ctx context.Context, id uint, image string) error {
	db := wrap_gorm.GetDB()
	updateInfo := UpdateOrderStatus{
		Status:        entities.DeliveredOrder,
		Image:         image,
		DeliveredDate: time.Now(),
	}
	err := db.Model(entities.Order{}).
		Where("id = ?", id).
		Updates(&updateInfo).
		Error
	return err
}

func (s *orderStorage) CountListByUserID(ctx context.Context, userID uint, input paging.ParamsInput) (total int64, err error) {
	db := wrap_gorm.GetDB()
	query := db.Model(entities.Order{}).Where("user_id = ?", userID)
	paging_query.SetCountListPagingQuery(&input, entities.Order{}.TableName(), query)
	err = query.Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (s *orderStorage) CountPreviewByUserID(ctx context.Context, userID uint, input paging.ParamsInput) (total int64, err error) {
	db := wrap_gorm.GetDB()
	query := db.Table(entities.Order{}.TableName()).
		Select("Order.*, JSON_ARRAYAGG(JSON_OBJECT('id', OrderItem.id, 'name', OrderItem.name, 'image', OrderItem.image,'price', OrderItem.price,'discount', OrderItem.discount,'quantity', OrderItem.quantity,'option', OrderItem.option)) as item").
		Joins("JOIN OrderItem on Order.id = OrderItem.order_id").
		Where("Order.user_id = ?", userID).
		Where("Order.deleted_at IS NULL").
		Where("OrderItem.deleted_at IS NULL").
		Group("Order.id")

	paging_query.SetCountListPagingQuery(&input, entities.Order{}.TableName(), query)
	err = query.Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (s *orderStorage) CountByProviderID(ctx context.Context, providerID uint, input paging.ParamsInput) (total int64, err error) {
	db := wrap_gorm.GetDB()
	query := db.Model(entities.Order{}).Where("provider_id = ?", providerID)
	paging_query.SetCountListPagingQuery(&input, entities.Order{}.TableName(), query)
	err = query.Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (s *orderStorage) CountPreviewByProviderID(ctx context.Context, providerID uint, input paging.ParamsInput) (total int64, err error) {
	db := wrap_gorm.GetDB()
	query := db.Table(entities.Order{}.TableName()).
		Select("Order.*, JSON_ARRAYAGG(JSON_OBJECT('id', OrderItem.id, 'name', OrderItem.name, 'image', OrderItem.image,'price', OrderItem.price,'discount', OrderItem.discount,'quantity', OrderItem.quantity,'option', OrderItem.option)) as item").
		Joins("JOIN OrderItem on Order.id = OrderItem.order_id").
		Where("Order.provider_id = ?", providerID).
		Where("Order.deleted_at IS NULL").
		Where("OrderItem.deleted_at IS NULL").
		Group("Order.id")

	paging_query.SetCountListPagingQuery(&input, entities.Order{}.TableName(), query)
	err = query.Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

//
func (s *orderStorage) CountList(ctx context.Context, input paging.ParamsInput) (total int64, err error) {
	db := wrap_gorm.GetDB()
	query := db.Model(entities.Order{})
	paging_query.SetCountListPagingQuery(&input, entities.Order{}.TableName(), query)
	err = query.Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (s *orderStorage) CountPreview(ctx context.Context, input paging.ParamsInput) (total int64, err error) {
	db := wrap_gorm.GetDB()
	query := db.Table(entities.Order{}.TableName()).
		Select("Order.*, JSON_ARRAYAGG(JSON_OBJECT('id', OrderItem.id, 'name', OrderItem.name, 'image', OrderItem.image,'price', OrderItem.price,'discount', OrderItem.discount,'quantity', OrderItem.quantity,'option', OrderItem.option)) as item").
		Joins("JOIN OrderItem on Order.id = OrderItem.order_id").
		Where("Order.deleted_at IS NULL").
		Where("OrderItem.deleted_at IS NULL").
		Group("Order.id")

	paging_query.SetCountListPagingQuery(&input, entities.Order{}.TableName(), query)
	err = query.Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func NewOrderStorage(orderItemSto order_item.Storage) order.Storage {
	return &orderStorage{orderItemSto: orderItemSto}
}
