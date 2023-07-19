package storage

import (
	"context"

	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
)

func (c commentStorage) DeleteCommentByID(ctx context.Context, commentID uint) error {
	db := wrap_gorm.GetDB()
	err := db.Table(entities.Comment{}.TableName()).
		Where("id = ?", commentID).
		Delete(&entities.Comment{}).
		Error
	if err != nil {
		return err
	}
	return nil
}
