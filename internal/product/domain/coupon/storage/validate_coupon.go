package storage

import (
	"context"
	"fmt"
	"github.com/eNViDAT0001/Thesis/Backend/external/product_quantities"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"gorm.io/gorm"
)

func (b couponStorage) ValidateCouponByProductIDs(ctx context.Context, couponCode string, productID uint) (entities.Coupon, error) {
	var result entities.Coupon
	db := wrap_gorm.GetDB()
	err := db.Table(entities.Coupon{}.TableName()).
		Joins("JOIN CouponDetail ON Coupon.id = CouponDetail.coupon_id").
		Where("Coupon.code = ?", couponCode).
		Where("CouponDetail.product_id = ?", productID).
		Where("CouponDetail.total > 0").
		Group("Coupon.id").
		First(&result).
		Error

	if err != nil {
		return result, err
	}

	var product entities.Product
	err = db.Model(entities.Product{}).Where("id = ?", productID).First(&product).Error
	if err != nil {
		return result, err
	}
	price := float64(product.Price) / ((100 - float64(product.Discount)) / 100.0)
	discountPrice := price * (100 - float64(result.Percent)/100.0)
	if discountPrice > result.Fixed {
		discountPrice = result.Fixed
	}
	result.Fixed = discountPrice

	return result, nil
}

func (b couponStorage) UseCouponByProductIDsWithGorm(ctx context.Context, db *gorm.DB, codes []string, productIDs []uint) error {
	quantityStore := product_quantities.GetCouponStore()
	store := map[uint]int{}

	for _, code := range codes {
		var couponDetail entities.CouponDetail
		err := db.Table(entities.CouponDetail{}.TableName()).
			Joins("JOIN Coupon ON Coupon.id = CouponDetail.coupon_id AND Coupon.code = ?", code).
			Joins("JOIN ProductPreview ON ProductPreview.id = CouponDetail.product_id AND ProductPreview.id IN ?", productIDs).
			First(&couponDetail).Error
		if err != nil {
			return err
		}

		store[couponDetail.ID] = 1
	}

	ok, invalidKey := quantityStore.Reduce(ctx, store)
	for invalidKey != 0 {
		var couponDetail entities.CouponDetail
		err := db.Table(entities.CouponDetail{}.TableName()).Where("id = ?", invalidKey).First(&couponDetail).Error
		if err != nil {
			return err
		}

		quantityStore.Add(ctx, invalidKey, couponDetail.Total)
		ok, invalidKey = quantityStore.Reduce(ctx, store)
	}

	if !ok {
		return fmt.Errorf("coupon is not have enough quantity")
	}

	for k, v := range store {
		err := db.Table(entities.CouponDetail{}.TableName()).
			Where("id = ?", k).
			UpdateColumn("total", gorm.Expr("total - ?", v)).
			Error
		if err != nil {
			return err
		}
	}
	return nil
}
