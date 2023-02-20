package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"

	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/address/entities"
)

func (a *addressStorage) GetProvinces(ctx context.Context) ([]entities.Province, error) {
	result := make([]entities.Province, 0)

	tableName := entities.Province{}.TableName()
	db := wrap_gorm.GetDB()
	//TODO: Need Join To Get AdministrativeRegion And AdministrativeUnit

	err := db.Table(tableName).
		Where("deleted_at IS NULL").
		Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}
