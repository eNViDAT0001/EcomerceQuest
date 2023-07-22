package storage

import Coupon "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/coupon"

type couponStorage struct {
}

func NewCouponStorage() Coupon.Storage {
	return &couponStorage{}
}
