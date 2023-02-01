package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"

	"github.com/eNViDAT0001/Thesis/Backend/internal/address/entities"
)

func (a *addressStorage) GetDistrictsWithProvinceCode(ctx context.Context, provinceCode string) ([]entities.District, error) {
	result := make([]entities.District, 0)
	tableName := entities.District{}.TableName()
	db := wrap_gorm.GetDB()
	//TODO: Need Join To Get AdministrativeRegion And AdministrativeUnit

	err := db.Table(tableName).
		Where("province_code = ?", provinceCode).
		Where("deleted_at IS NULL").
		Find(&result).Error

	if err != nil {

		return nil, err
	}

	return result, nil
}
