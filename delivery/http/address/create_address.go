package address

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/address/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s addressHandler) CreateAddress() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.CreateAddressReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}
		userID, _ := strconv.Atoi(cc.Param("user_id"))
		input.UserID = strconv.Itoa(userID)

		err := s.addressUC.CreateAddress(newCtx, &input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Create Address success")
	}
}
