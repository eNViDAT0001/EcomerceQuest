package product

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func (s *productHandler) ListProductWithCategoryID() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		paginator, err := paging_query.GetPagingParams(cc.Context, entities.Product{})
		if err != nil {
			cc.BadRequest(err)
			return
		}

		categoryID, err := strconv.Atoi(cc.Param("category_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		categories, err := s.categoryUC.GetCategoryChildrenByCategoryID(newCtx, uint(categoryID))
		if err != nil {
			cc.NoContent(err)
			return
		}

		categoryIDs := make([]uint, 0)
		for _, category := range categories {
			categoryIDs = append(categoryIDs, category.ID)
		}
		categoryIDs = append(categoryIDs, uint(categoryID))
		inputRepo := io.ListProductInput{
			CategoryIDs: categoryIDs,
			Paging:      paginator,
		}
		products, total, err := s.productUC.ListProduct(newCtx, inputRepo)
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
