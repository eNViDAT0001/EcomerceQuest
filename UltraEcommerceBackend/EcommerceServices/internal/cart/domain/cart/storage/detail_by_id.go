package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/cart/entities"
	"gorm.io/gorm"
)

func (c cartStorage) GetDetailByID(ctx context.Context, cartID uint) (entities.CartDetail, error) {
	result := entities.CartDetail{}
	db := wrap_gorm.GetDB()
	query := db.Table(entities.CartDetail{}.TableName()).Where("id = ?", cartID)

	query = query.Scan(&result)
	err := query.Error
	if err != nil {
		return result, err
	}
	if query.RowsAffected < 1 {
		return result, gorm.ErrRecordNotFound
	}

	return result, err
}
