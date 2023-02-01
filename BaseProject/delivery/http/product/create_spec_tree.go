package product

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/product/convert"
	ioHandler "github.com/eNViDAT0001/Thesis/Backend/delivery/http/product/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *productHandler) CreateSpecificationTree() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		productID, err := strconv.Atoi(cc.Param("product_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		var input ioHandler.ProductSpecificationsCreateTreeReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		if len(input.Specification) < 1 {
			cc.ResponseError(request.NewBadRequestError("Specification", nil, "Empty specifications"))
		}

		for i := range input.Specification {
			input.Specification[i].Specification.ProductID = uint(productID)
			for j := 0; j < len(input.Specification[i].Options); j++ {
				input.Specification[i].Options[j].ProductID = uint(productID)
			}
		}

		inputUC, err := convert.CreateSpecificationArrayReqToCreateSpecificationForm(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		err = s.productUC.CreateSpecificationTree(newCtx, inputUC)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Create Product Specification Success")
	}
}
