package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
)

func (s productStorage) DeleteProductByID(ctx context.Context, ID uint) error {
	db := wrap_gorm.GetDB()
	err := db.Table(entities.Product{}.TableName()).
		Where("id = ?", ID).
		Delete(&entities.Product{}).
		Error
	if err != nil {
		return err
	}
	return nil
}
