package product

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/grpc/grpc_base"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	proto "github.com/eNViDAT0001/Thesis/Backend/thesis_proto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

var userStorage = map[uint][]uint{}

func (s *productHandler) ListRecommendedProductsPreview() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		paginator, err := paging_query.GetPagingParams(cc.Context, entities.ProductPreview{})
		if err != nil {
			cc.ResponseError(err)
			return
		}

		userID, err := strconv.Atoi(cc.Param("user_id"))
		if err != nil || userID == 0 {
			inputRepo := io.ListProductInput{
				Paging: paginator,
			}
			products, total, err := s.productUC.ListProductsPreview(newCtx, inputRepo)
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					cc.NoContent(err)
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
			return
		}

		recommendedIDs, ok := userStorage[uint(userID)]
		if paginator.Marker < 1 || !ok {
			productIDs, err := grpc_base.GetServices().RecommenderService.
				LisRecommendedProductIDsByUserID(newCtx, &proto.RecommendReq{UserId: int32(userID)})
			userStorage[uint(userID)] = productIDs

			inputRepo := io.ListRecommendedProductInput{
				RecommendedProductIDs: productIDs,
				Paging:                paginator,
			}
			products, total, err := s.productUC.ListRecommendedProductsPreview(newCtx, inputRepo)
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					cc.NoContent(err)
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
			return
		}

		inputRepo := io.ListRecommendedProductInput{
			RecommendedProductIDs: recommendedIDs,
			Paging:                paginator,
		}
		products, total, err := s.productUC.ListRecommendedProductsPreview(newCtx, inputRepo)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent(err)
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
