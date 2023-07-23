package storage

import (
	coupon "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/coupon"
)

type couponStorage struct {
}

func NewCouponStorage() coupon.Storage {
	return &couponStorage{}
}
