package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/category/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

func (c categoryStorage) DeleteCategoryNodeByID(ctx context.Context, categoryID uint, input *ioSto.CategoryForm) error {
	db := wrap_gorm.GetDB()

	var childrenIDs []uint
	childrenIDs = append(childrenIDs, categoryID)

	children, err := c.GetCategoryChildrenTreeWithCategoryID(ctx, categoryID)
	for _, child := range children {
		childrenIDs = append(childrenIDs, child.ID)
	}

	err = db.Model(entities.Category{}).Where("id IN ?", childrenIDs).Delete(&entities.Category{}).Error

	if err != nil {
		return err
	}

	return nil
}
