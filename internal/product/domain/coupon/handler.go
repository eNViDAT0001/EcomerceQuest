package Coupon

import (
	"github.com/gin-gonic/gin"
)

type HttpHandler interface {
	CreateCoupon() func(*gin.Context)
	GetCouponByID() func(*gin.Context)
	UpdateCoupon() func(*gin.Context)
	DeleteCouponByIDs() func(*gin.Context)
	ListCoupon() func(*gin.Context)
	ListProductByCouponID() func(*gin.Context)
	ListProductPreviewByCouponID() func(*gin.Context)
	ListProductNotInCouponID() func(*gin.Context)
	ListProductPreviewNotInCouponID() func(*gin.Context)

	ValidateCouponByProductID() func(*gin.Context)
	UseCoupon() func(*gin.Context)
}
