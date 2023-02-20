package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/entities"
)

func (s productStorage) GetProductDescriptionsByProductID(ctx context.Context, productID uint) ([]entities.ProductDescriptions, error) {
	var result []entities.ProductDescriptions
	db := wrap_gorm.GetDB()
	err := db.Model(entities.ProductDescriptions{}).
		Where("product_id = ?", productID).
		Find(&result).Error
	if err != nil {
		return result, err
	}
	return result, err
}
