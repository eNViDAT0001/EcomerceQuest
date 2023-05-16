package order_item

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	ListByOrderID() func(ctx *gin.Context)
}
