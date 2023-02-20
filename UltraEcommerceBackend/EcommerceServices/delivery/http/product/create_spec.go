package product

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/product/convert"
	ioHandler "github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/product/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func (s *productHandler) CreateSpecification() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		productID, err := strconv.Atoi(cc.Param("product_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		var input ioHandler.SpecificationCreate
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		input.Specification.ProductID = uint(productID)
		for i := 0; i < len(input.Options); i++ {
			input.Options[i].ProductID = uint(productID)
		}

		inputUC, err := convert.CreateSpecificationReqToCreateSpecificationForm(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		specID, err := s.productUC.CreateSpecification(newCtx, inputUC)
		if err != nil {
			if err == gorm.ErrRegistered {
				cc.Conflict(err)
			}
			cc.ResponseError(err)
			return
		}

		result := map[string]interface{}{
			"SpecificationID": specID,
		}
		cc.Ok(result)
	}
}
