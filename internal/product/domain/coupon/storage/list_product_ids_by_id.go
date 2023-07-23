package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
)

func (b couponStorage) ListProductIDsByCouponID(ctx context.Context, couponID uint, filter paging.ParamsInput) ([]uint, error) {
	result := make([]uint, 0)

	db := wrap_gorm.GetDB()

	query := db.Table(entities.Coupon{}.TableName()).
		Select("DISTINCT CouponDetail.product_id").
		Joins("JOIN CouponDetail ON CouponDetail.coupon_id = Coupon.id").
		Where("Coupon.id = ?", couponID).
		Where("Coupon.deleted_at IS NULL")

	paging_query.SetPagingQuery(&filter, entities.Coupon{}.TableName(), query)

	err := query.Find(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil
}
