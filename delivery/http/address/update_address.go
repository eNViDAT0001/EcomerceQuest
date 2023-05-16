package address

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/address/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s addressHandler) UpdateAddress() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.UpdateAddressReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		id, err := strconv.Atoi(cc.Param("address_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		err = s.addressUC.UpdateAddress(newCtx, uint(id), input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Update Address success")
	}
}
