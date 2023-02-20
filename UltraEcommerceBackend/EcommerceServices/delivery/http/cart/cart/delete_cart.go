package cart

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *cartHandler) DeleteCart() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		cartID, err := strconv.Atoi(cc.Param("cart_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}
		userID, err := strconv.Atoi(cc.Param("user_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		err = s.cartUC.DeleteCart(newCtx, uint(userID), uint(cartID))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Delete Product success")
	}
}
