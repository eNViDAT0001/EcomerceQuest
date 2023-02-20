package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/entities"
)

func (s productStorage) UpdateProductDescriptions(ctx context.Context, descriptionsID uint, descriptions io.ProductDescriptionsUpdateForm) error {
	db := wrap_gorm.GetDB()
	err := db.Table(entities.ProductDescriptions{}.TableName()).
		Where("id = ?", descriptionsID).
		Updates(&descriptions).Error
	if err != nil {
		return err
	}
	return nil
}
