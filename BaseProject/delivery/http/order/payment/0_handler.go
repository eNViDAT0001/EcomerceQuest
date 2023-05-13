package order_items

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/payment"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/payment/storage/io"
	"github.com/gin-gonic/gin"
)

type paymentHandler struct {
	paymentUC payment.UseCase
}

func (s *paymentHandler) GetPaymentByID() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		cc := request.FromContext(ctx)
		newCtx := context.Background()

		paymentID := cc.Param("payment_id")
		result, err := s.paymentUC.GetPaymentByID(newCtx, paymentID)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok(result)
	}
}
func (s *paymentHandler) GetPaypalPaymentByID() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		cc := request.FromContext(ctx)
		//newCtx := context.Background()
		//
		//paymentID := cc.Param("payment_id")
		//paypalClient, err := paypal.GetClient()
		//if err != nil {
		//	cc.ResponseError(err)
		//}
		cc.Ok("result")
	}
}

func (s *paymentHandler) CreatePayment() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		cc := request.FromContext(ctx)
		newCtx := context.Background()

		var input io.CreatePaymentForm
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		payment, err := s.paymentUC.CreatePayment(newCtx, input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok(payment)
	}
}

func NewPaymentHandler(paymentUC payment.UseCase) payment.HttpHandler {
	return &paymentHandler{paymentUC: paymentUC}
}
