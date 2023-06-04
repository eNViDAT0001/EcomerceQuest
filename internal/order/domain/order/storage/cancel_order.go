package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/product_quantities"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
	entities2 "github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"gorm.io/gorm"
)

func (s orderStorage) CancelOrder(ctx context.Context, orderID uint) error {
	db := wrap_gorm.GetDB()
	err := db.Model(entities.Order{}).Where("id = ?", orderID).Update("status", string(entities.CancelOrder)).Error
	if err != nil {
		return err
	}

	items, err := s.orderItemSto.ListByOrderID(ctx, orderID)
	if err != nil {
		return err
	}

	quantityStore := product_quantities.GetQuantityStore()
	for _, v := range items {
		err = db.Table(entities2.ProductOption{}.TableName()).
			Where("id = ?", v.ProductOptionID).
			UpdateColumn("quantity", gorm.Expr("quantity + ?", v.Quantity)).
			Error
		if err != nil {
			return err
		}

		quantityStore.Add(ctx, v.ProductOptionID, v.Quantity)
	}
	return nil
}
