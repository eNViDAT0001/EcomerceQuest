package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
)

func (s orderStorage) DeleteOrder(ctx context.Context, orderID uint) error {
	db := wrap_gorm.GetDB()
	query := db.Begin()
	err := query.Model(entities.Order{}).Where("id = ?", orderID).Delete(&entities.Order{}).Error
	if err != nil {
		query.Rollback()
		return err
	}
	err = query.Model(entities.OrderItem{}).Where("order_id = ?", orderID).Delete(&entities.OrderItem{}).Error
	if err != nil {
		query.Rollback()
		return err
	}

	query.Commit()
	return nil
}
func (s orderStorage) DeleteOrders(ctx context.Context, ids []uint) error {
	db := wrap_gorm.GetDB()
	query := db.Begin()
	err := query.Model(entities.Order{}).Where("id IN ?", ids).Delete(&entities.Order{}).Error
	if err != nil {
		query.Rollback()
		return err
	}
	err = query.Model(entities.OrderItem{}).Where("order_id IN ?", ids).Delete(&entities.OrderItem{}).Error
	if err != nil {
		query.Rollback()
		return err
	}

	query.Commit()
	return nil
}
