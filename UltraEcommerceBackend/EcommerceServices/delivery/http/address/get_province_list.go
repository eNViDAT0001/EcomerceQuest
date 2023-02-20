package address

import (
	"context"

	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/address/domain/address/usecase/convert"
	"github.com/gin-gonic/gin"
)

func (s addressHandler) GetProvinces() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		provinces, err := s.addressUC.GetProvinces(newCtx)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		if len(provinces) == 0 {
			cc.NoContent()
			return
		}

		result, err := convert.ProvinceArrayEntityToProvinceArrayRes(provinces)

		cc.Ok(result)
	}
}
