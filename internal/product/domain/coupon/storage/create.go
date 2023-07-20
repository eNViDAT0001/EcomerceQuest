package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
)

func (b couponStorage) CreateCoupon(ctx context.Context, input CouponCreateForm, productIDs []uint) (BannerID uint, err error) {
	db := wrap_gorm.GetDB()
	err = db.Table(entities.Banner{}.TableName()).Create(&input).Error
	if err != nil {
		return 0, err
	}

	productStorage := make([]entities.Coupon, 0)
	for _, id := range productIDs {
		couponDetail := entities.CouponDetail{
			CouponID:  input.ID,
			ProductID: id,
		}
		productStorage = append(productStorage, bannerDetail)
	}

	err = db.Table(entities.Coupon{}.TableName()).Create(&productStorage).Error
	if err != nil {
		return 0, err
	}
	return input.ID, nil
}
