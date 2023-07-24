package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

func (c categoryStorage) DeleteCategoryNodeByID(ctx context.Context, categoryID uint) error {
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
