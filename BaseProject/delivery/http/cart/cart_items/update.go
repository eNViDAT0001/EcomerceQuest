package cart_items

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/cart/cart_items/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *cartItemHandler) UpdateCartItem() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.CartItemUpdateReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		cartID, _ := strconv.Atoi(cc.Param("cart_id"))
		itemID, _ := strconv.Atoi(cc.Param("cart_item_id"))

		err := s.cartItemUC.UpdateCartItem(newCtx, uint(itemID), uint(cartID), input.Quantity)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Upsert CartItem Success")
	}
}
