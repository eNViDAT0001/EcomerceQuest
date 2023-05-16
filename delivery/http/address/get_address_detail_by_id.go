package address

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s addressHandler) GetAddressDetailByID() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		addressID, _ := strconv.Atoi(cc.Param("address_id"))

		addresses, err := s.addressUC.GetAddressDetailByID(newCtx, uint(addressID))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent()
				return
			}
			cc.ResponseError(err)
			return
		}

		cc.Ok(addresses)
	}
}
