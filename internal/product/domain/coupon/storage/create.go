package storage

import (
	"context"
	"errors"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/coupon/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (b couponStorage) CreateCoupon(ctx context.Context, input io.CouponCreateForm, products []io.CouponDetailCreateForm) (CouponID uint, err error) {
	database := wrap_gorm.GetDB()

	err = database.Transaction(func(db *gorm.DB) error {
		var coupon entities.Coupon
		err = db.Table(entities.Coupon{}.TableName()).First(&coupon).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		if coupon.Code == input.Code {
			return errors.New("coupon code already exists")
		}

		err = db.Table(entities.Coupon{}.TableName()).Create(&input).Error
		if err != nil {
			return err
		}

		for i, _ := range products {
			products[i].CouponID = input.ID
		}

		err = db.Table(entities.CouponDetail{}.TableName()).Clauses(clause.OnConflict{
			DoNothing: true,
		}).Create(&products).Error

		if err != nil {
			return err
		}
		return nil
	})

	return input.ID, err
}
