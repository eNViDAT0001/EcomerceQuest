package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/entities"
)

func (s providerStorage) ListProvider(ctx context.Context) ([]entities.Provider, error) {
	result := make([]entities.Provider, 0)
	db := wrap_gorm.GetDB()
	err := db.Model(entities.Provider{}).Find(&result).Error
	if err != nil {

		return result, err
	}
	return result, nil
}
