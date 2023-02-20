package product

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/product/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *productHandler) CreateProductMedia() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.ProductMedia
		if err := cc.ShouldBind(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		productID, err := strconv.Atoi(cc.Param("product_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		err = s.productUC.CreateProductMediaWithFiles(newCtx, uint(productID), input.Files)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		cc.Ok("Create Product Media Success")
	}
}
