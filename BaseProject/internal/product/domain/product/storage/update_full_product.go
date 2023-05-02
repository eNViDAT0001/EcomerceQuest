package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"gorm.io/gorm"
)

func (s productStorage) UpdateFullProduct(ctx context.Context, productID uint, product io.ProductFullUpdateForm) error {
	db := wrap_gorm.GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		productUpdateForm := io.ProductUpdateForm{
			Name:       product.Name,
			Price:      product.Price,
			Discount:   product.Discount,
			CategoryID: product.CategoryID,
		}
		err := tx.Model(entities.Product{}).
			Where("id = ?", productID).
			Updates(&productUpdateForm).Error
		if err != nil {
			return err
		}

		for _, option := range product.Options {
			err = tx.Model(entities.ProductOption{}).
				Where("id = ?", option.ID).
				Where("product_id = ?", productID).
				Updates(&option.Option).Error
			if err != nil {
				return err
			}
		}

		for _, description := range product.Descriptions {
			err = tx.Table(entities.ProductDescriptions{}.TableName()).
				Where("id = ?", description.ID).
				Where("product_id = ?", productID).
				Updates(&description.Description).Error
			if err != nil {
				return err
			}
		}
		for _, media := range product.Media {
			err = tx.Model(entities.ProductMedia{}).
				Where("id = ?", media.ID).
				Where("product_id = ?", productID).
				Update("media_path", media.MediaPath).Error
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
