package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/coupon/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (b couponStorage) UpdateCoupon(ctx context.Context, couponID uint, input io.CouponUpdateForm, productsIN []io.CouponDetailCreateForm, productIDsOUT []uint) error {
	database := wrap_gorm.GetDB()

	err := database.Transaction(func(db *gorm.DB) error {
		err := db.Model(entities.Coupon{}).Where("id = ?", couponID).Updates(&input).Error
		if err != nil {
			return err
		}
		for i, _ := range productsIN {
			productsIN[i].CouponID = couponID
		}
		if len(productsIN) > 0 {
			err = db.Table(entities.CouponDetail{}.TableName()).Clauses(clause.OnConflict{
				UpdateAll: true,
			}).Create(&productsIN).Error
			if err != nil {
				db.Rollback()
				return err
			}
		}

		if len(productIDsOUT) > 0 {
			err = db.Table(entities.CouponDetail{}.TableName()).
				Where("product_id IN ?", productIDsOUT).
				Where("coupon_id = ?", couponID).
				Delete(&entities.CouponDetail{}).
				Error

			if err != nil {
				db.Rollback()
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
