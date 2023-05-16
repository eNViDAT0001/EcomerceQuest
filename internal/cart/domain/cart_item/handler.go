package cart_item

import (
	"github.com/gin-gonic/gin"
)

type HttpHandler interface {
	UpsertCartItem() func(ctx *gin.Context)
	UpdateCartItem() func(ctx *gin.Context)
	DeleteCartItem() func(ctx *gin.Context)
}
