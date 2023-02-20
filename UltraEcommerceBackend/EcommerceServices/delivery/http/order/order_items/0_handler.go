package order_items

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/order/domain/order_item"
	"github.com/gin-gonic/gin"
	"strconv"
)

type orderItemHandler struct {
	orderItemUC order_item.UseCase
}

func (s *orderItemHandler) ListByOrderID() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		cc := request.FromContext(ctx)
		newCtx := context.Background()

		orderID, err := strconv.Atoi(cc.Param("order_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		orders, err := s.orderItemUC.ListByOrderID(newCtx, uint(orderID))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok(orders)
	}
}

func NewOrderItemHandler(orderItemUC order_item.UseCase) order_item.HttpHandler {
	return &orderItemHandler{orderItemUC: orderItemUC}
}
