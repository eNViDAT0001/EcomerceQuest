package product

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s *productHandler) GetProductMediaByProductID() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		productID, err := strconv.Atoi(cc.Param("product_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		media, err := s.productUC.ListProductMediaByProductID(newCtx, uint(productID))
		if err != nil {
			cc.ResponseError(err)
			return
		}
		if len(media) < 1 {
			cc.NoContent()
			return
		}
		cc.Ok(media)
	}
}
