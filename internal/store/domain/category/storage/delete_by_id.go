package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

func (c categoryStorage) DeleteCategoryByID(ctx context.Context, categoryID uint, parentID *uint) error {
	db := wrap_gorm.GetDB()

	err := db.Table(entities.Category{}.TableName()).
		Where("id = ?", categoryID).
		Delete(&entities.Category{}).
		Error
	if err != nil {
		return err
	}
	//Nếu có parentID thì nối children vào parentID
	if parentID != nil && *parentID != 0 {
		err := db.Model(entities.Category{}).Where("category_parent_id = ?", categoryID).Update("category_parent_id", parentID).Error
		if err != nil {
			return err
		}
		return nil
	}
	// Nếu không có ParentID thì nối children vào nil - Cập nhật children lên root
	err = db.Model(entities.Category{}).Where("category_parent_id = ?", categoryID).Update("category_parent_id", nil).Error
	if err != nil {
		return err
	}

	return nil
}
