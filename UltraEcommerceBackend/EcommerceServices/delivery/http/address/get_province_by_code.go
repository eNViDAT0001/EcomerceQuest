package address

import (
	"context"

	"gorm.io/gorm"

	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/address/domain/address/usecase/convert"
	"github.com/gin-gonic/gin"
)

func (s addressHandler) GetProvinceByCode() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		provinceCode := cc.Param("province_code")

		provinces, err := s.addressUC.GetProvinceByCode(newCtx, provinceCode)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		result, err := convert.ProvinceEntityToProvinceRes(&provinces)
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
