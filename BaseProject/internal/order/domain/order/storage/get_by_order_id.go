package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
)

func (s orderStorage) GetByOrderID(ctx context.Context, orderID uint) (entities.Order, error) {
	var result entities.Order
	db := wrap_gorm.GetDB()
	err := db.Model(entities.Order{}).Where("id = ?", orderID).Find(&result).Error
	if err != nil {
		return result, err
	}
	return result, nil
}
