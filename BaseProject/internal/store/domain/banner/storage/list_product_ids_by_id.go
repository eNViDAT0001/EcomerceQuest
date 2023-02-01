package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

func (b bannerStorage) ListProductIDsByBannerID(ctx context.Context, bannerID uint, filter paging.ParamsInput) ([]uint, error) {
	result := make([]uint, 0)

	db := wrap_gorm.GetDB()

	query := db.Table(entities.Banner{}.TableName()).
		Select("BannerDetail.product_id").
		Joins("JOIN BannerDetail ON BannerDetail.banner_id = Banner.id").
		Where("Banner.id = ?", bannerID).
		Where("Banner.deleted_at IS NULL")

	paging_query.SetPagingQuery(&filter, entities.Banner{}.TableName(), query)

	err := query.Find(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil
}
