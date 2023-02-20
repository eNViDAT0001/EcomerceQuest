package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/entities"
)

func (c categoryStorage) ListCategories(ctx context.Context) ([]entities.Category, error) {
	var result []entities.Category
	db := wrap_gorm.GetDB()
	err := db.Model(entities.Category{}).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, err
}
