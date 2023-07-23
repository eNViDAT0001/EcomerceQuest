package coupon

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/coupon/convert"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/coupon/io"

	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s couponHandler) CreateCoupon() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.CouponCreateReq
		if err := cc.ShouldBind(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		inputSto, err := convert.CreateReqToCreateCouponInput(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		productSto, err := convert.CreateReqToCreateCouponProductsInput(input.Products)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		bannerID, err := s.couponUC.CreateCoupon(newCtx, inputSto, productSto)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		result := map[string]interface{}{
			"BannerID": bannerID,
		}
		cc.Ok(result)
	}
}
