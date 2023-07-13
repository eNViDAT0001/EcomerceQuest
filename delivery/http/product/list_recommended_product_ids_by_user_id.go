package product

import (
	"context"
	"encoding/json"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/grpc/grpc_base"
	"github.com/eNViDAT0001/Thesis/Backend/external/cache"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	proto "github.com/eNViDAT0001/Thesis/Backend/thesis_proto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func (s *productHandler) ListRecommendedProductsIds() func(ctx *gin.Context) {
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
			if paginator.Filter.GetSort() == nil {
				paginator.Filter = paginator.Filter.CloneWithSort([]string{"rating_DESC"})
			}
			(*paginator.Filter.GetSort())["rating"] = "DESC"

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

		recommendedIDs, ok := getRecommendedIDs(newCtx, uint(userID))

		if paginator.Marker < 1 || !ok {
			productIDs, err := grpc_base.GetServices().RecommenderService.
				LisRecommendedProductIDsByUserID(newCtx, &proto.RecommendReq{UserId: int32(userID)})
			if err != nil {
				if paginator.Filter.GetSort() == nil {
					paginator.Filter = paginator.Filter.CloneWithSort([]string{"rating_DESC"})
				}
				(*paginator.Filter.GetSort())["rating"] = "DESC"

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
			saveRecommendedIDs(newCtx, uint(userID), productIDs)

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

func getRecommendedIDs(ctx context.Context, id uint) (productIDs []uint, exist bool) {
	val, err := cache.GetRedis().Get(ctx, "recommended_"+strconv.Itoa(int(id)))
	if err != nil {
		log.Println(err)
		return nil, false
	}

	if err = json.Unmarshal([]byte(val), &productIDs); err != nil {
		log.Println(err)
		return nil, false
	}

	return productIDs, true
}

func saveRecommendedIDs(ctx context.Context, id uint, recommendedIDs []uint) {
	err := cache.GetRedis().SetDefault(ctx, "recommended_"+strconv.Itoa(int(id)), recommendedIDs)
	if err != nil {
		log.Println(err)
	}
}
