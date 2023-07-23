package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	io "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/coupon/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"gorm.io/gorm"
)

func (b couponStorage) GetCouponByID(ctx context.Context, couponID uint) (io.CouponDetail, error) {
	var result io.CouponDetail
	db := wrap_gorm.GetDB()
	query := db.Table(entities.Coupon{}.TableName()).
		Select("Coupon.*, IF(COUNT(Product.id) = 0, NULL, JSON_ARRAYAGG(JSON_OBJECT( 'id', Product.id, 'name', Product.name, 'total', CouponDetail.total))) AS products").
		Joins("LEFT JOIN CouponDetail ON CouponDetail.coupon_id = Coupon.id").
		Joins("JOIN Product ON Product.id = CouponDetail.product_id").
		Where("Coupon.id = ?", couponID).
		Group("Coupon.id").
		Scan(&result)

	err := query.Error
	if err != nil {
		return result, err
	}
	if query.RowsAffected < 1 {
		return result, gorm.ErrRecordNotFound
	}

	return result, nil
}
