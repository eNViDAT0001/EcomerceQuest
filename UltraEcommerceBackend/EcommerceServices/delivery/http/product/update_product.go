package product

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/product/convert"
	ioHandler "github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/product/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *productHandler) UpdateProduct() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		productID, err := strconv.Atoi(cc.Param("product_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		var input ioHandler.ProductUpdateReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		inputRepo, err := convert.UpdateProductReqToUpdateProductForm(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		err = s.productUC.UpdateProduct(newCtx, uint(productID), inputRepo)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Update Product Success")
	}
}
