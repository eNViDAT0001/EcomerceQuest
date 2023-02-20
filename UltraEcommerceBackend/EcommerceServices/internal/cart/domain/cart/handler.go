package cart

import (
	"github.com/gin-gonic/gin"
)

type HttpHandler interface {
	ListCartByUserID() func(ctx *gin.Context)
	DeleteCart() func(ctx *gin.Context)
}
