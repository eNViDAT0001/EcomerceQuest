package storage

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/coupon"
)

type couponStorage struct {
}

func NewCouponStorage() coupon.Storage {
	return &couponStorage{}
}
