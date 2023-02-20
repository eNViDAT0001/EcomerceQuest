package cart_items

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/cart/cart_items/convert"
	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/cart/cart_items/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *cartItemHandler) UpsertCartItem() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.CartItemCreateReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		productID, _ := strconv.Atoi(cc.Param("product_id"))
		providerID, _ := strconv.Atoi(cc.Param("provider_id"))
		userID, _ := strconv.Atoi(cc.Param("user_id"))

		input.UserID = uint(userID)
		input.ProductID = uint(productID)

		inputSto, err := convert.CreateReqToCreateCartItemsFormInput(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		err = s.cartItemUC.UpsertCartItem(newCtx, uint(providerID), inputSto)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Upsert CartItem Success")
	}
}
