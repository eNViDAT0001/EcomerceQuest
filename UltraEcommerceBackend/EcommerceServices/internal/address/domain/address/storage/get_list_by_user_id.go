package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"

	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/address/domain/address/storage/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/address/entities"
)

func (a *addressStorage) GetAddressesWithUserID(ctx context.Context, userID uint) ([]io.AddressPreview, error) {
	result := make([]io.AddressPreview, 0)
	tableName := entities.Address{}.TableName()
	db := wrap_gorm.GetDB()

	err := db.Table(tableName).
		Select("Address.id, "+
			"Address.user_id, "+
			"Address.name, "+
			"Address.gender, "+
			"Address.phone, "+
			"Provinces.name AS province, "+
			"Districts.name AS district, "+
			"Wards.name AS ward, "+
			"Address.street").
		Joins("JOIN Districts ON Districts.code = Address.district_code").
		Joins("JOIN Wards ON Wards.code = Address.ward_code").
		Joins("JOIN Provinces ON Provinces.code = Address.province_code").
		Where("user_id = ?", userID).
		Where("Address.deleted_at IS NULL").
		Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}
