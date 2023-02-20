package product

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/product/convert"
	ioHandler "github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/product/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
)

func (s *productHandler) CreateDescriptions() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input ioHandler.ProductDescriptionsWithFileCreateReq
		if err := cc.ShouldBind(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		productID, _ := strconv.Atoi(cc.Param("product_id"))
		if input.ProductID != uint(productID) {
			cc.Conflict(request.NewConflictError("Descriptions.ProductID", input.ProductID, "ProductID and Descriptions.ProductID does not match"))
			return
		}

		inputUC, err := convert.CreateDescriptionsReqToCreateDescriptionsForm(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		descriptionsID, err := s.productUC.CreateDescriptions(newCtx, inputUC)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		result := map[string]interface{}{
			"DescriptionsID": descriptionsID,
		}
		cc.Ok(result)
	}
}
