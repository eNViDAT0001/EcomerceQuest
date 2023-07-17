package order

import (
	"context"

	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s *orderHandler) ListReport() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		paginator, err := paging_query.GetPagingParams(cc.Context, entities.Order{})
		if err != nil {
			cc.ResponseError(err)
			return
		}

		orders, _, err := s.orderUC.ListOrderReport(newCtx, paginator)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent()
				return
			}
			cc.ResponseError(err)
			return
		}

		cc.Ok(orders)
	}
}
