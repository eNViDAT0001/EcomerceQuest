package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"

	"github.com/eNViDAT0001/Thesis/Backend/internal/address/entities"
)

func (a *addressStorage) GetAddressesWithUserID(ctx context.Context, userID uint) ([]entities.Address, error) {
	result := make([]entities.Address, 0)
	db := wrap_gorm.GetDB()

	err := db.Model(entities.Address{}).
		Where("user_id = ?", userID).
		Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}
