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
func (s productStorage) DeleteElementByIDs(ctx context.Context, ID uint, descriptionsIDs []uint, mediaIDs []uint, optionsIDs []uint) error {
	db := wrap_gorm.GetDB()
	err := db.Table(entities.ProductDescriptions{}.TableName()).
		Where("id IN ?", descriptionsIDs).
		Where("product_id = ?", ID).
		Delete(&entities.ProductDescriptions{}).
		Error
	if err != nil {
		return err
	}
	err = db.Table(entities.ProductMedia{}.TableName()).
		Where("id IN ?", mediaIDs).
		Where("product_id = ?", ID).
		Delete(&entities.ProductMedia{}).
		Error
	if err != nil {
		return err
	}
	err = db.Table(entities.ProductOption{}.TableName()).
		Where("id IN ?", optionsIDs).
		Where("product_id = ?", ID).
		Delete(&entities.ProductOption{}).
		Error
	if err != nil {
		return err
	}
	return nil
}
