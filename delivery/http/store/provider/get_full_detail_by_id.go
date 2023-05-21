package provider

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s providerHandler) GetProviderFullDetailByID() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		providerID, err := strconv.Atoi(cc.Param("provider_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		result, err := s.providerUC.GetProviderFullDetailByID(newCtx, uint(providerID))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok(result)
	}
}
