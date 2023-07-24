package coupon

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/coupon/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s couponHandler) ValidateCouponByProductID() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.ValidateCouponReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		result, err := s.couponUC.ValidateCouponByProductIDs(newCtx, input.Code, input.ProductIDs)
		if err != nil && err != gorm.ErrRecordNotFound {
			cc.ResponseError(err)
			return
		}

		cc.Ok(result)
	}
}
