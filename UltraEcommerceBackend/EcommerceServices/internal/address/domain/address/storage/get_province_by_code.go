package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"

	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/address/entities"
)

func (a *addressStorage) GetProvinceByCode(ctx context.Context, provinceCode string) (entities.Province, error) {
	var result entities.Province
	tableName := entities.Province{}.TableName()
	db := wrap_gorm.GetDB()
	//TODO: Need Join To Get AdministrativeRegion And AdministrativeUnit
	err := db.Table(tableName).
		Where("code = ?", provinceCode).
		Where("deleted_at IS NULL").
		First(&result).Error

	if err != nil {
		return result, err
	}

	return result, nil
}
