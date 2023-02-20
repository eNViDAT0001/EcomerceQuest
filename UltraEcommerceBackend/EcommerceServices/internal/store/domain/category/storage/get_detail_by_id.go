package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/entities"
)

func (c categoryStorage) GetCategoryDetailByID(ctx context.Context, categoryID uint) (entities.Category, error) {
	var result entities.Category
	db := wrap_gorm.GetDB()

	err := db.Model(entities.Category{}).Where("id = ?", categoryID).First(result).Error

	if err != nil {
		return result, err
	}

	return result, nil
}
