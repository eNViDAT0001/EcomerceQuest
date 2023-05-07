package order

import (
	"context"
	"errors"
	io2 "github.com/eNViDAT0001/Thesis/Backend/delivery/http/order/order/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *orderHandler) UpdateOrderStatus() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io2.UpdateOrderStatusReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}
		orderID, err := strconv.Atoi(cc.Param("order_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		if input.Status == entities.DeliveredOrder {
			if input.Image == "" {
				cc.BadRequest(errors.New("image is required"))
				return
			}
			err = s.orderUC.UpdateDeliveredOrderStatus(newCtx, uint(orderID), input.Image)
			if err != nil {
				cc.ResponseError(err)
				return
			}
			cc.Ok("Update Status success")
			return
		}

		err = s.orderUC.UpdateOrderStatus(newCtx, uint(orderID), input.Status)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Update Status success")
	}
}
