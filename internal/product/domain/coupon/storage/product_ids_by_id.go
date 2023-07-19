package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	entities2 "github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

func (b couponStorage) ProductIDsByBannerID(ctx context.Context, bannerID uint) ([]uint, error) {
	result := make([]uint, 0)

	db := wrap_gorm.GetDB()

	query := db.Table(entities.Banner{}.TableName()).
		Select("DISTINCT BannerDetail.product_id").
		Joins("JOIN BannerDetail ON BannerDetail.banner_id = Banner.id").
		Where("Banner.id = ?", bannerID).
		Where("Banner.deleted_at IS NULL")

	err := query.Find(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil
}
func (b couponStorage) ProductIDsByNotInBannerID(ctx context.Context, bannerID uint) ([]uint, error) {
	result := make([]uint, 0)
	productIDs, err := b.ProductIDsByBannerID(ctx, bannerID)
	if err != nil {
		return result, err
	}
	db := wrap_gorm.GetDB()

	query := db.Table(entities2.Product{}.TableName()).
		Select("Product.id").
		Where("Product.id NOT IN (?)", productIDs).
		Where("Product.deleted_at IS NULL")

	err = query.Find(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil
}
