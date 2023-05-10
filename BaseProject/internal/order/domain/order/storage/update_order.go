package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
)

func (s orderStorage) UpdateOrder(ctx context.Context, orderID uint, input io.UpdateOrderForm) error {
	db := wrap_gorm.GetDB()
	err := db.Model(entities.Order{}).Where("id = ?", orderID).Updates(&input).Error
	if err != nil {
		return err
	}
	return nil
}
