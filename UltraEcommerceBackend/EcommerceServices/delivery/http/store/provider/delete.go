package provider

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s providerHandler) DeleteProviderByID() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		providerID, err := strconv.Atoi(cc.Param("provider_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}
		ids := []uint{uint(providerID)}
		err = s.providerUC.DeleteProviderByIDs(newCtx, ids)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Delete Provider Success")
	}
}
