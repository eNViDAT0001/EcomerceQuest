package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
	"gorm.io/gorm"
)

func (c categoryStorage) RecoverCategoryNodeByID(ctx context.Context, categoryID uint) error {
	database := wrap_gorm.GetDB()

	err := database.Transaction(func(db *gorm.DB) error {
		var childrenIDs []uint

		children, err := c.GetCategoryChildrenTreeWithCategoryID(ctx, categoryID)
		for _, child := range children {
			childrenIDs = append(childrenIDs, child.ID)
		}

		category := entities.Category{}
		err = db.Table(entities.Category{}.TableName()).
			Where("id = ?", categoryID).
			First(&category).
			Error
		if err != nil {
			return err
		}

		if category.CategoryParentID == nil {
			*category.CategoryParentID = 0
		}

		category = entities.Category{}
		err = db.Table(entities.Category{}.TableName()).
			Where("id = ?", category.CategoryParentID).
			First(&category).
			Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				err = db.Table(entities.Category{}.TableName()).
					Where("id = ?", categoryID).
					Update("category_parent_id", nil).
					Update("deleted_at", nil).
					Error
				if err != nil {
					return err
				}
				err = db.Table(entities.Category{}.TableName()).
					Where("category_parent_id IN ?", childrenIDs).
					Update("deleted_at", nil).
					Error

				return err
			}
			return err
		}

		err = db.Table(entities.Category{}.TableName()).
			Where("id = ?", categoryID).
			Update("deleted_at", nil).
			Error
		if err != nil {
			return err
		}
		err = db.Table(entities.Category{}.TableName()).
			Where("id IN ?", childrenIDs).
			Update("deleted_at", nil).
			Error
		return err
	})
	return err
}
