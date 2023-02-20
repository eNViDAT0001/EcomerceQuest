package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
	ioSto "github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/domain/category/storage/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/entities"
)

func (c categoryStorage) CreateCategory(ctx context.Context, input *ioSto.CategoryForm) error {
	db := wrap_gorm.GetDB()

	err := db.Table(entities.Category{}.TableName()).Create(input).Error

	if err != nil {
		return err
	}

	return nil
}
