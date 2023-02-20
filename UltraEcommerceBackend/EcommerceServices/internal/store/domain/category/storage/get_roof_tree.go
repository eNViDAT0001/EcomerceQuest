package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/entities"
)

func (c categoryStorage) GetCategoryRoofTree(ctx context.Context) ([]entities.Category, error) {
	result := make([]entities.Category, 0)

	db := wrap_gorm.GetDB()

	err := db.Model(entities.Category{}).Where("category_parent_id IS NULL").Find(&result).Error

	if err != nil {

		return result, err
	}

	return result, nil
}
