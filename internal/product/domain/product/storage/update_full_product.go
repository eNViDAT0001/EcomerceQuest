package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/enum"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"gorm.io/gorm"
)

func (s productStorage) UpdateFullProduct(ctx context.Context, productID uint, product io.ProductFullUpdateForm) error {
	db := wrap_gorm.GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		productUpdateForm := io.ProductUpdateForm{
			Name:              product.Name,
			Price:             product.Price,
			Discount:          product.Discount,
			CategoryID:        product.CategoryID,
			ShortDescriptions: product.ShortDescriptions,
			Height:            product.Height,
			Width:             product.Width,
			Length:            product.Length,
			Weight:            product.Weight,
		}
		err := tx.Model(entities.Product{}).
			Where("id = ?", productID).
			Updates(&productUpdateForm).Error
		if err != nil {
			return err
		}

		for _, option := range product.Options {
			if option.ID != 0 && option.Option.SpecificationID != 0 {
				return request.NewConflictError("specification_id", option.Option.SpecificationID, "specification_id and id cannot be set at the same time")
			}
			if option.Option.SpecificationID == 0 && option.ID != 0 {
				err = tx.Model(entities.ProductOption{}).
					Where("id = ?", option.ID).
					Where("product_id = ?", productID).
					Updates(&option.Option).Error
				if err != nil {
					return err
				}
				continue
			}

			if option.ID == 0 && option.Option.SpecificationID == 0 {
				return request.NewConflictError("id", option.ID, "specification_id and id cannot be empty at the same time")
			}
			option.Option.ProductID = productID
			err = tx.Table(entities.ProductOption{}.TableName()).Create(&option.Option).Error
			if err != nil {
				return err
			}
		}

		for _, description := range product.Descriptions {
			if description.ID != 0 {
				err = tx.Table(entities.ProductDescriptions{}.TableName()).
					Where("id = ?", description.ID).
					Where("product_id = ?", productID).
					Updates(&description.Description).Error
				if err != nil {
					return err
				}
				continue
			}
			description.Description.ProductID = productID
			err = tx.Table(entities.ProductDescriptions{}.TableName()).
				Create(&description.Description).Error
			if err != nil {
				return err
			}

		}
		for _, media := range product.Media {
			if media.ID != 0 {
				err = tx.Model(entities.ProductMedia{}).
					Where("id = ?", media.ID).
					Where("product_id = ?", productID).
					Update("media_path", media.Media.MediaPath).Error
				if err != nil {
					return err
				}
				continue
			}
			media.Media.ProductID = productID
			media.Media.MediaType = string(enum.MediaImage)
			err = tx.Table(entities.ProductMedia{}.TableName()).
				Create(&media.Media).Error
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
