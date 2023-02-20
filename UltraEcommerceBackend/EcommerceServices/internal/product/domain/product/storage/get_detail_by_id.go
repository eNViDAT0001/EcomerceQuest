package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/entities"
)

func (s productStorage) GetProductDetailByID(ctx context.Context, productID uint) (entities.Product, error) {
	var result entities.Product
	db := wrap_gorm.GetDB()
	err := db.Model(entities.Product{}).
		Where("id = ?", productID).
		First(&result).Error
	if err != nil {
		return result, err
	}
	return result, err
}
