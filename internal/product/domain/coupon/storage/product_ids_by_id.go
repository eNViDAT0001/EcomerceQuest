package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
)

func (b couponStorage) ProductIDsByCouponID(ctx context.Context, couponID uint) ([]uint, error) {
	result := make([]uint, 0)

	db := wrap_gorm.GetDB()

	query := db.Table(entities.Coupon{}.TableName()).
		Select("DISTINCT CouponDetail.product_id").
		Joins("JOIN CouponDetail ON CouponDetail.coupon_id = Coupon.id").
		Joins("JOIN Product ON Product.id = CouponDetail.product_id AND Product.deleted_at IS NULL").
		Where("Coupon.id = ?", couponID)

	err := query.Find(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil
}
func (b couponStorage) ProductIDsByNotInCouponID(ctx context.Context, couponID uint) ([]uint, error) {
	result := make([]uint, 0)
	productIDs, err := b.ProductIDsByCouponID(ctx, couponID)
	if err != nil {
		return result, err
	}
	db := wrap_gorm.GetDB()

	query := db.Table(entities.Product{}.TableName()).
		Select("Product.id").
		Where("Product.id NOT IN (?)", productIDs).
		Where("Product.deleted_at IS NULL")

	err = query.Find(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil
}
