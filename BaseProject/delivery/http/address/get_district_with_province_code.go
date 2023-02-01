package address

import (
	"context"

	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s addressHandler) GetDistrictsWithProvinceCode() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		provinceCode := cc.Param("province_code")

		districts, err := s.addressUC.GetDistrictsWithProvinceCode(newCtx, provinceCode)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		if len(districts) == 0 {
			cc.NoContent()
			return
		}
		cc.Ok(districts)
	}
}
