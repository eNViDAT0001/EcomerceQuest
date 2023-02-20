package product

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/product/convert"
	ioHandler "github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/product/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
)

func (s *productHandler) CreateProductOptions() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input ioHandler.ProductOptionsCreateReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		productID, _ := strconv.Atoi(cc.Param("product_id"))
		for _, otp := range input.Options {
			if otp.ProductID != uint(productID) {
				cc.Conflict(request.NewConflictError("Option.ProductID", otp.ProductID, "ProductID and Options.ProductID does not match"))
				return
			}
		}

		inputUC, err := convert.CreateOptionsReqToCreateOptionsForm(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		err = s.productUC.CreateProductOptions(newCtx, inputUC)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Create Product Options Success")
	}
}
