package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"

	"github.com/eNViDAT0001/Thesis/Backend/internal/address/entities"
)

func (a *addressStorage) GetAddressDetailByID(ctx context.Context, addressID uint) (entities.Address, error) {
	var result entities.Address
	db := wrap_gorm.GetDB()

	err := db.Model(entities.Address{}).
		Where("Address.id = ?", addressID).
		First(&result).Error

	if err != nil {
		return result, err
	}

	return result, nil
}
