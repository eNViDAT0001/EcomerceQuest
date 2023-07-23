package coupon

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func (s couponHandler) GetCouponByID() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		id, err := strconv.Atoi(cc.Param("coupon_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		result, err := s.couponUC.GetCouponByID(newCtx, uint(id))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent()
				return
			}
			cc.ResponseError(err)
			return
		}

		cc.Ok(result)
	}
}
