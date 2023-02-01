package provider

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func (s providerHandler) GetProviderByID() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		providerID, err := strconv.Atoi(cc.Param("provider_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		result, err := s.providerUC.GetProviderByID(newCtx, uint(providerID))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent()
				return
			}
			cc.ResponseError(err)
			return
		}

		cc.Ok(result)
	}
}
