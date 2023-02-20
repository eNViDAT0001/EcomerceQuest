package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/entities"
)

func (s productStorage) CreateProductMedia(ctx context.Context, media []io.CreateProductMedia) error {
	db := wrap_gorm.GetDB()
	err := db.Table(entities.ProductMedia{}.TableName()).Create(&media).Error
	if err != nil {
		return err
	}
	return nil
}
