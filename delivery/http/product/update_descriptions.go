package product

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/product/convert"
	ioHandler "github.com/eNViDAT0001/Thesis/Backend/delivery/http/product/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s *productHandler) UpdateDescriptions() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		descriptionsID, err := strconv.Atoi(cc.Param("description_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}
		var input ioHandler.ProductDescriptionsWithFileUpdateReq
		if err := cc.ShouldBind(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		inputUC, err := convert.UpdateDescriptionsReqToUpdateDescriptionsForm(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		err = s.productUC.UpdateProductDescriptions(newCtx, uint(descriptionsID), inputUC)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Update Descriptions Update")
	}
}
