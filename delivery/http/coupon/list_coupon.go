package coupon

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	paging_query "github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s couponHandler) ListCoupon() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		paginator, err := paging_query.GetPagingParams(cc.Context, entities.Coupon{})
		if err != nil {
			cc.ResponseError(err)
			return
		}

		coupons, total, err := s.couponUC.ListCoupon(newCtx, paginator)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent()
				return
			}
			cc.ResponseError(err)
			return
		}

		paginator.Total = int(total)
		if paginator.Type == paging.CursorPaging && len(coupons) > 0 {
			paginator.Marker = int(coupons[len(coupons)-1].ID)
		}

		cc.OkPaging(paginator, coupons)
	}
}
