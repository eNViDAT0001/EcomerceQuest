package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/cart/entities"
	"gorm.io/gorm"
)

func (c cartItemStorage) DeleteCartItem(ctx context.Context, cartID uint, itemID uint) error {
	db := wrap_gorm.GetDB()
	query := db.Table(entities.CartItem{}.TableName()).
		Where("id = ?", itemID).
		Where("cart_id = ?", cartID)

	query = query.Delete(&entities.CartItem{})
	err := query.Error
	if err != nil {
		return err
	}
	if query.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
