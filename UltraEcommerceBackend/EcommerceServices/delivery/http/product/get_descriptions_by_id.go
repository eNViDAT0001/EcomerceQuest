package product

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *productHandler) GetProductDescriptionsByProductID() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		productID, err := strconv.Atoi(cc.Param("product_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		descriptions, err := s.productUC.GetProductDescriptionsByProductID(newCtx, uint(productID))
		if err != nil {
			cc.ResponseError(err)
			return
		}
		if len(descriptions) < 1 {
			cc.NoContent()
			return
		}
		cc.Ok(descriptions)
	}
}
