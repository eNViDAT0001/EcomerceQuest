package address

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s addressHandler) GetAddressesWithUserID() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		userID, _ := strconv.Atoi(cc.Param("user_id"))

		addresses, err := s.addressUC.GetAddressesWithUserID(newCtx, uint(userID))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		if len(addresses) == 0 {
			cc.NoContent()
			return
		}

		cc.Ok(addresses)
	}
}
