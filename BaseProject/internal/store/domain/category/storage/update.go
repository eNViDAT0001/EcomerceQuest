package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/category/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

func (c categoryStorage) UpdateCategory(ctx context.Context, categoryID uint, input *ioSto.CategoryForm) error {
	db := wrap_gorm.GetDB()

	err := db.Model(entities.Category{}).Where("id = ?", categoryID).Updates(input).Error
	if input.CategoryParentID != nil && *input.CategoryParentID == 0 {
		err = db.Model(entities.Category{}).Where("id = ?", categoryID).Update("category_parent_id", nil).Error
		if err != nil {
			return err
		}
	}

	if err != nil {
		return err
	}

	return nil
}
