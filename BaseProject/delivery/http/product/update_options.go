package product

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/product/convert"
	ioHandler "github.com/eNViDAT0001/Thesis/Backend/delivery/http/product/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s *productHandler) UpdateProductOptions() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		productID, err := strconv.Atoi(cc.Param("product_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		var input ioHandler.ProductOptionsUpdateReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		for _, otp := range input.Options {
			if otp.ProductID != uint(productID) {
				cc.Conflict(request.NewConflictError("Option.ProductID", otp.ProductID, "ProductID and Options.ProductID does not match"))
				return
			}
		}

		inputRepo, err := convert.UpdateOptionsReqToUpdateOptionsForm(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		err = s.productUC.UpdateProductOptions(newCtx, inputRepo)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		cc.Ok("Update Product Options Success")
	}

}
