package address

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/address/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
)

func (s addressHandler) DeleteAddressByIDs() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.DeleteAddressReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}
		userID, _ := strconv.Atoi(cc.Param("user_id"))
		err := s.addressUC.DeleteAddressByIDs(newCtx, uint(userID), input.IDs)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Delete Address Success")
	}
}
