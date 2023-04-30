package order

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type orderHandler struct {
	orderUC order.UseCase
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

func NewOrderHandler(orderUC order.UseCase) order.HttpHandler {
	return &orderHandler{orderUC: orderUC}
}
