package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/address/domain/address/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/address/entities"
)

func (a *addressStorage) CreateAddress(ctx context.Context, input *io.FullAddressForm) error {
	tableName := entities.Address{}.TableName()
	db := wrap_gorm.GetDB()

	err := db.Table(tableName).Create(&input).Error

	if err != nil {
		return err
	}

	return nil
}
