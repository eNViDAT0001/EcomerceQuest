package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
)

func (s productStorage) GetProductSpecificationRoofByProductID(ctx context.Context, productID uint) ([]entities.ProductSpecification, error) {

	result := make([]entities.ProductSpecification, 0)
	db := wrap_gorm.GetDB()
	err := db.Model(entities.ProductSpecification{}).
		Where("product_id = ?", productID).
		Find(&result).Error
	if err != nil {
		return result, err
	}
	return result, err
}
