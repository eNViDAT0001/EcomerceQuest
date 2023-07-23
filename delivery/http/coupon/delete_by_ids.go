package coupon

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/coupon/io"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s couponHandler) DeleteCouponByIDs() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.CouponIDsReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		userID, err := strconv.Atoi(cc.Param("user_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		err = s.couponUC.DeleteCouponByIDs(newCtx, uint(userID), input.IDs)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		cc.Ok("Delete Coupon Success")
	}
}
