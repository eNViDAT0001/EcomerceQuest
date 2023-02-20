package provider

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
)

func (s providerHandler) ListProvider() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		result, err := s.providerUC.ListProvider(newCtx)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		if len(result) == 0 {
			cc.NoContent()
			return
		}
		cc.Ok(result)
	}
}
