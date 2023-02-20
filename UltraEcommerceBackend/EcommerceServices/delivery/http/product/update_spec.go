package product

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/product/convert"
	ioHandler "github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/product/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
)

func (s *productHandler) UpdateProductSpecification() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		specID, err := strconv.Atoi(cc.Param("specification_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}
		productID, err := strconv.Atoi(cc.Param("product_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		var input ioHandler.ProductSpecificationUpdateReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}
		if input.ProductID != uint(productID) {
			cc.Conflict(request.NewConflictError("Specification.ProductID", input.ProductID, "ProductID and Specification.ProductID does not match"))
			return
		}

		inputRepo, err := convert.UpdateSpecificationReqToUpdateSpecificationForm(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		err = s.productUC.UpdateProductSpecification(newCtx, uint(specID), inputRepo)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		cc.Ok("Update Product Options Success")
	}
}
