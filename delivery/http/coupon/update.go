package coupon

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/coupon/convert"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/coupon/io"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s couponHandler) UpdateCoupon() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.CouponUpdateReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}
		couponID, err := strconv.Atoi(cc.Param("coupon_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}
		userID, err := strconv.Atoi(cc.Param("user_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}
		input.UserID = uint(userID)

		inputSto, products, err := convert.UpdateReqToUpdateCouponInput(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		err = s.couponUC.UpdateCoupon(newCtx, uint(couponID), inputSto, products, input.ProductIDsOUT)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		result := map[string]interface{}{
			"CouponID": couponID,
		}

		cc.Ok(result)
	}
}
