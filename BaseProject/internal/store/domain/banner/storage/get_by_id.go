package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/banner/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
	"gorm.io/gorm"
)

func (b bannerStorage) GetBannerByID(ctx context.Context, bannerID uint) (io.BannerDetail, error) {
	var result io.BannerDetail
	db := wrap_gorm.GetDB()
	query := db.Table(entities.Banner{}.TableName()).
		Select("Banner.*, IF(COUNT(Product.id) = 0, NULL, JSON_ARRAYAGG(JSON_OBJECT( 'id', Product.id, 'name', Product.name))) AS products").
		Joins("LEFT JOIN BannerDetail ON BannerDetail.banner_id = Banner.id").
		Joins("JOIN Product ON Product.id = BannerDetail.product_id").
		Where("Banner.id = ?", bannerID).
		Where("Banner.deleted_at IS NULL").
		Group("Banner.id").
		Scan(&result)

	err := query.Error
	if err != nil {
		return result, err
	}
	if query.RowsAffected < 1 {
		return result, gorm.ErrRecordNotFound
	}

	return result, nil
}
