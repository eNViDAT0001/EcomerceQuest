package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
)

func (b couponStorage) ValidateCouponByProductIDs(ctx context.Context, couponCode string, productID uint) (entities.Coupon, error) {
	var result entities.Coupon
	db := wrap_gorm.GetDB()
	err := db.Table(entities.Coupon{}.TableName()).
		Joins("JOIN CouponDetail ON CouponDetail.id = CouponDetail.coupon_id").
		Where("CouponDetail.product_id = ?", productID).
		Where("CouponDetail.code = ?", couponCode).
		Where("CouponDetail.total > 0").
		Where("CouponDetail.deleted_at IS NULL").
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

//
//func (b couponStorage) UseCouponByProductIDsWithGorm(ctx context.Context, db *gorm.DB, coupon []entities.Coupon, productIDs []uint) (entities.Coupon, error) {
//	quantityStore := product_quantities.GetCouponStore()
//	store := map[uint]int{}
//
//	for i, v := range productIDs {
//
//		store[productIDs] += v.Quantity
//	}
//	ok, invalidKey := quantityStore.Reduce(ctx, store)
//
//	for invalidKey != 0 {
//		var option entities2.ProductOption
//		err = query.Table(entities2.ProductOption{}.TableName()).
//			Where("id = ?", invalidKey).First(&option).Error
//		if err != nil {
//			query.Rollback()
//			return nil, err
//		}
//
//		quantityStore.Add(ctx, invalidKey, option.Quantity)
//		ok, invalidKey = quantityStore.Reduce(ctx, store)
//	}
//
//	if !ok {
//		query.Rollback()
//		return nil, fmt.Errorf("product is not have enough quantity")
//	}
//
//	return result, nil
//}
