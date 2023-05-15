package product

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func (s *productHandler) ListProductPreviewWithBannerID() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		paginator, err := paging_query.GetPagingParams(cc.Context, entities.Product{})
		if err != nil {
			cc.BadRequest(err)
			return
		}

		bannerID, err := strconv.Atoi(cc.Param("banner_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		inputRepo := io.ListProductInput{
			BannerID: uint(bannerID),
			Paging:   paginator,
		}
		products, total, err := s.productUC.ListProductPreviewByBannerID(newCtx, inputRepo)
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
func (s *productHandler) ListProductPreviewNotInBannerID() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		paginator, err := paging_query.GetPagingParams(cc.Context, entities.Product{})
		if err != nil {
			cc.BadRequest(err)
			return
		}

		bannerID, err := strconv.Atoi(cc.Param("banner_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		inputRepo := io.ListProductInput{
			BannerID: uint(bannerID),
			Paging:   paginator,
		}
		products, total, err := s.productUC.ListProductPreviewNotInBannerID(newCtx, inputRepo)
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
