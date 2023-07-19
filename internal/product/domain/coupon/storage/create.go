package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/banner/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

func (b couponStorage) CreateBanner(ctx context.Context, input io.BannerCreateForm, productIDs []uint) (BannerID uint, err error) {
	db := wrap_gorm.GetDB()
	err = db.Table(entities.Banner{}.TableName()).Create(&input).Error
	if err != nil {
		return 0, err
	}

	productStorage := make([]entities.BannerDetail, 0)
	for _, id := range productIDs {
		bannerDetail := entities.BannerDetail{
			BannerID:  input.ID,
			ProductID: id,
		}
		productStorage = append(productStorage, bannerDetail)
	}

	err = db.Table(entities.BannerDetail{}.TableName()).Create(&productStorage).Error
	if err != nil {
		return 0, err
	}
	return input.ID, nil
}
