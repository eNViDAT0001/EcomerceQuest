package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/address/entities"
)

func (a *addressStorage) DeleteAddressByIDs(ctx context.Context, userID uint, addressIDs []uint) error {
	tableName := entities.Address{}.TableName()
	db := wrap_gorm.GetDB()

	err := db.Table(tableName).
		Where("id IN ?", addressIDs).
		Where("user_id = ?", userID).
		Delete(&entities.Address{}).Error

	if err != nil {
		return err
	}

	return nil
}
