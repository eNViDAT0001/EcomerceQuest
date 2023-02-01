package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
)

func (s *orderItemStorage) ListByOrderID(ctx context.Context, orderID uint) ([]entities.OrderItem, error) {
	result := make([]entities.OrderItem, 0)
	db := wrap_gorm.GetDB()
	err := db.Model(entities.OrderItem{}).Where("order_id = ?", orderID).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
