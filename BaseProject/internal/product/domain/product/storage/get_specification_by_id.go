package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
)

func (s productStorage) GetSpecificationByID(ctx context.Context, id uint) (entities.ProductSpecification, error) {
	var result entities.ProductSpecification
	db := wrap_gorm.GetDB()
	err := db.Model(entities.ProductSpecification{}).
		Where("id = ?", id).
		First(&result).Error
	if err != nil {
		return result, err
	}
	return result, err
}
