package order

import (
	"context"
	io2 "github.com/eNViDAT0001/Thesis/Backend/delivery/http/order/order/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/domain/user"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type orderHandler struct {
	orderUC  order.UseCase
	useUC    user.UseCase
	smtpUC   smtp.UseCase
	notifyUC notification.UseCase
}

func (s *orderHandler) RemoveInvalidOrder() error {
	ctx := context.Background()
	unPayOrders, unConfirmedOrders, err := s.orderUC.ListInvalidOrder(ctx)
	if err != nil {
		return err
	}

	for _, v := range unPayOrders {
		err = s.orderUC.CancelOrder(ctx, v.ID)
		if err != nil {
			return err
		}
	}

	delivered := true
	for _, v := range unConfirmedOrders {
		err = s.orderUC.UpdateOrder(ctx, v.ID, io.UpdateOrderForm{
			VerifyDelivered: &delivered,
		})
		
		if err != nil {
			return err
		}
	}

	return err
}

func (s *orderHandler) UpdateOrder() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.UpdateOrderForm
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}
		orderID, err := strconv.Atoi(cc.Param("order_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		err = s.orderUC.UpdateOrder(newCtx, uint(orderID), input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Update Status success")
	}
}
func (s *orderHandler) UpdateOrderPayment() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io2.UpdateOrderPaymentReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}
		inputUC := io.UpdateOrderPaymentForm{
			PaymentID:  input.PaymentID,
			PaymentURL: input.PaymentURL,
		}
		err := s.orderUC.UpdateOrderPayment(newCtx, input.OrderIDs, inputUC)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Update Order Payment success")
	}
}

func (s *orderHandler) VerifyDeliveredStatus() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		orderID, err := strconv.Atoi(cc.Param("order_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}
		userID, err := strconv.Atoi(cc.Param("user_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		err = s.orderUC.VerifyDeliveredOrder(newCtx, uint(orderID), uint(userID))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Verify Status success")
	}
}

func (s *orderHandler) List() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		paginator, err := paging_query.GetPagingParams(cc.Context, entities.Order{})
		if err != nil {
			cc.ResponseError(err)
			return
		}

		orders, total, err := s.orderUC.List(newCtx, paginator)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent()
				return
			}
			cc.ResponseError(err)
			return
		}

		paginator.Total = int(total)
		if paginator.Type == paging.CursorPaging && len(orders) > 0 {
			paginator.Marker = int(orders[len(orders)-1].ID)
		}

		cc.OkPaging(paginator, orders)
	}
}

func (s *orderHandler) ListPreview() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		paginator, err := paging_query.GetPagingParams(cc.Context, entities.Order{})
		if err != nil {
			cc.ResponseError(err)
			return
		}

		orders, total, err := s.orderUC.ListPreview(newCtx, paginator)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent()
				return
			}
			cc.ResponseError(err)
			return
		}

		paginator.Total = int(total)
		if paginator.Type == paging.CursorPaging && len(orders) > 0 {
			paginator.Marker = int(orders[len(orders)-1].ID)
		}

		cc.OkPaging(paginator, orders)
	}
}

func NewOrderHandler(orderUC order.UseCase, smtpUC smtp.UseCase, useUC user.UseCase, notifyUC notification.UseCase) order.HttpHandler {
	return &orderHandler{orderUC: orderUC, smtpUC: smtpUC, useUC: useUC, notifyUC: notifyUC}
}
