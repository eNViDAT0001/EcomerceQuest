package coupon

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func (s couponHandler) ListProductPreviewByCouponID() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		couponID, err := strconv.Atoi(cc.Param("coupon_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		paginator, err := paging_query.GetPagingParams(cc.Context, entities.Product{})
		if err != nil {
			cc.ResponseError(err)
			return
		}

		products, total, err := s.couponUC.ListProductPreviewByCouponID(newCtx, uint(couponID), paginator)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent()
				return
			}
			cc.ResponseError(err)
			return
		}

		paginator.Total = int(total)
		if paginator.Type == paging.CursorPaging && len(products) > 0 {
			paginator.Marker = int(products[len(products)-1].ID)
		}

		cc.OkPaging(paginator, products)
	}
}
func (s couponHandler) ListProductPreviewNotInCouponID() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		bannerID, err := strconv.Atoi(cc.Param("coupon_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		paginator, err := paging_query.GetPagingParams(cc.Context, entities.Product{})
		if err != nil {
			cc.ResponseError(err)
			return
		}

		products, total, err := s.couponUC.ListProductPreviewNotInCouponID(newCtx, uint(bannerID), paginator)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent()
				return
			}
			cc.ResponseError(err)
			return
		}

		paginator.Total = int(total)
		if paginator.Type == paging.CursorPaging && len(products) > 0 {
			paginator.Marker = int(products[len(products)-1].ID)
		}

		cc.OkPaging(paginator, products)
	}
}
