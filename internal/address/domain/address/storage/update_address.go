package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"

	"github.com/eNViDAT0001/Thesis/Backend/internal/address/domain/address/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/address/entities"
)

func (a *addressStorage) UpdateAddress(ctx context.Context, addressID uint, input *io.AddressForm) error {
	tableName := entities.Address{}.TableName()
	db := wrap_gorm.GetDB()

	err := db.Table(tableName).
		Where("id = ?", addressID).
		Where("deleted_at IS NULL").
		Updates(&input).Error

	if err != nil {

		return err
	}

	return nil
}
