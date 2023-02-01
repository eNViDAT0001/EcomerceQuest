package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"

	"github.com/eNViDAT0001/Thesis/Backend/internal/address/domain/address/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/address/entities"
)

func (a *addressStorage) GetAddressDetailByID(ctx context.Context, addressID uint) (io.AddressDetail, error) {
	var result io.AddressDetail
	tableName := entities.Address{}.TableName()
	db := wrap_gorm.GetDB()

	err := db.Table(tableName).
		Select("Address.id, "+
			"Address.user_id, "+
			"Address.name, "+
			"Address.gender, "+
			"Address.phone, "+
			"Address.district_code, "+
			"Address.ward_code, "+
			"Address.province_code, "+
			"Provinces.name AS province, "+
			"Districts.name AS district, "+
			"Wards.name AS ward, "+
			"Address.street").
		Joins("JOIN Districts ON Districts.code = Address.district_code").
		Joins("JOIN Wards ON Wards.code = Address.ward_code").
		Joins("JOIN Provinces ON Provinces.code = Address.province_code").
		Where("Address.id = ?", addressID).
		Where("Address.deleted_at IS NULL").
		First(&result).Error

	if err != nil {
		return result, err
	}

	return result, nil
}
