package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
)

func (s productStorage) UpdateProduct(ctx context.Context, productID uint, product io.ProductUpdateForm) error {
	db := wrap_gorm.GetDB()
	err := db.Model(entities.Product{}).
		Where("id = ?", productID).
		Updates(&product).Error
	if err != nil {
		return err
	}
	return nil
}
