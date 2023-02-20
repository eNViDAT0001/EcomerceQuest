package provider

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/store/provider/convert"
	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/store/provider/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
)

func (s providerHandler) UpdateProvider() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.UpdateProviderReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		inputSto, err := convert.ProviderUpdateReqToProviderUpdateForm(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		providerID, err := strconv.Atoi(cc.Param("provider_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		err = s.providerUC.UpdateProvider(newCtx, uint(providerID), inputSto)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Update Provider Success")
	}
}
