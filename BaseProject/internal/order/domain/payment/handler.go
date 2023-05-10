package payment

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	GetPaymentByID() func(ctx *gin.Context)
	CreatePayment() func(ctx *gin.Context)
}
