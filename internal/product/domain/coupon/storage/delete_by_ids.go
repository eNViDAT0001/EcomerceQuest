package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"gorm.io/gorm"
)

func (b couponStorage) DeleteCouponByIDs(ctx context.Context, userID uint, couponIDs []uint) error {
	database := wrap_gorm.GetDB()

	err := database.Transaction(func(db *gorm.DB) error {
		err := db.Table(entities.CouponDetail{}.TableName()).
			Where("coupon_id IN ?", couponIDs).Delete(&entities.CouponDetail{}).Error
		if err != nil {
			return err
		}

		err = db.Table(entities.Coupon{}.TableName()).
			Where("id IN ?", couponIDs).
			Where("user_id = ?", userID).
			Delete(&entities.Coupon{}).Error
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
