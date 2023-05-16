package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
)

func (s productStorage) ListProductMediaByProductID(ctx context.Context, productID uint) ([]entities.ProductMedia, error) {
	result := make([]entities.ProductMedia, 0)
	db := wrap_gorm.GetDB()
	err := db.Model(entities.ProductMedia{}).
		Where("product_id = ?", productID).
		Find(&result).Error
	if err != nil {
		return result, err
	}
	return result, err
}
