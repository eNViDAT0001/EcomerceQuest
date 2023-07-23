package coupon

import (
	coupon "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/coupon"
)

type couponHandler struct {
	couponUC coupon.UseCase
}

func NewBannerHandler(couponUC coupon.UseCase) coupon.HttpHandler {
	return &couponHandler{couponUC: couponUC}
}
