package product

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s *productHandler) GetSpecificationTreeByProductID() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		productID, err := strconv.Atoi(cc.Param("product_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		specTree, err := s.productUC.GetSpecificationTreeByProductID(newCtx, uint(productID))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent()
				return
			}

			cc.ResponseError(err)
			return
		}
		cc.Ok(specTree)
	}
}
