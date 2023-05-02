package product

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *productHandler) UpdateFullProduct() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		productID, err := strconv.Atoi(cc.Param("product_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		var input ioSto.ProductFullUpdateForm
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		err = s.productUC.UpdateFullProduct(newCtx, uint(productID), input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Update Product Success")
	}
}
