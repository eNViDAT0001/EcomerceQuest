package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/cart/domain/cart_item/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/cart/entities"
	"gorm.io/gorm"
)

func (c cartItemStorage) UpsertCartItem(ctx context.Context, providerID uint, input io.CartItemCreateForm) error {
	db := wrap_gorm.GetDB()
	var item entities.CartItem
	query := db.Begin()
	err := query.Model(entities.CartItem{}).
		Where("product_id = ?", input.ProductID).
		Where("user_id = ?", input.UserID).
		Where("product_option_id = ?", input.ProductOptionID).
		First(&item).
		Error

	if err == nil {
		errUpdate := query.Table(entities.CartItem{}.TableName()).
			Where("id = ?", item.ID).
			UpdateColumn("quantity", gorm.Expr("quantity + ?", input.Quantity)).Error
		if errUpdate != nil {
			query.Rollback()
			return errUpdate
		}
		query.Commit()
		return nil
	}

	if err != gorm.ErrRecordNotFound {
		query.Rollback()
		return err
	}
	// TODO: Reduce these code using Upsert instead of a transaction
	var cart entities.Cart
	err = query.Model(entities.Cart{}).
		Where("user_id = ?", input.UserID).
		Where("provider_id = ?", providerID).
		First(&cart).Error

	if err == nil {
		input.CartID = cart.ID
		err = query.Table(entities.CartItem{}.TableName()).
			Create(&input).
			Error
		if err != nil {
			query.Rollback()
			return err
		}
		query.Commit()
		return nil
	}

	if err != gorm.ErrRecordNotFound {
		query.Rollback()
		return err
	}

	newCart := entities.Cart{
		ProviderID: providerID,
		UserID:     input.UserID,
	}
	err = query.Table(entities.Cart{}.TableName()).Create(&newCart).Error
	if err != nil {
		query.Rollback()
		return err
	}

	newCartItem := entities.CartItem{
		CartID:          newCart.ID,
		UserID:          input.UserID,
		ProductID:       input.ProductID,
		ProductOptionID: input.ProductOptionID,
		Quantity:        input.Quantity,
	}
	err = query.Table(entities.CartItem{}.TableName()).
		Create(&newCartItem).
		Error

	if err != nil {
		query.Rollback()
		return err
	}

	query.Commit()

	return nil
}
