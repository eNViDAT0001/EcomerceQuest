package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
)

type orderStorage struct {
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

func NewOrderStorage() order.Storage {
	return &orderStorage{}
}
