package address

import (
	"context"

	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
)

func (s addressHandler) GetWardWithDistrictCode() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		districtCode := cc.Param("district_code")

		wards, err := s.addressUC.GetWardWithDistrictCode(newCtx, districtCode)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		if len(wards) == 0 {
			cc.NoContent()
			return
		}
		cc.Ok(wards)
	}
}
