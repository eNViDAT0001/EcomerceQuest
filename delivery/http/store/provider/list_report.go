package provider

import (
	"context"

	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s providerHandler) ListProviderReport() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		paginator, err := paging_query.GetPagingParams(cc.Context, entities.Provider{})
		if err != nil {
			cc.ResponseError(err)
			return
		}

		providers, _, err := s.providerUC.ListProviderQuantity(newCtx, paginator)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent()
				return
			}
			cc.ResponseError(err)
			return
		}

		cc.Ok(providers)
	}
}
