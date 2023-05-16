package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/cart/entities"
)

func (c cartStorage) DeleteCart(ctx context.Context, userID uint, cartID uint) error {
	db := wrap_gorm.GetDB()
	err := db.Table(entities.Cart{}.TableName()).
		Where("user_id = ?", userID).
		Where("id = ?", cartID).
		Delete(&entities.Cart{}).
		Error
	if err != nil {
		return nil
	}
	err = db.Table(entities.CartItem{}.TableName()).
		Where("cart_id = ?", cartID).
		Delete(&entities.CartItem{}).
		Error
	if err != nil {
		return nil
	}
	return nil
}
