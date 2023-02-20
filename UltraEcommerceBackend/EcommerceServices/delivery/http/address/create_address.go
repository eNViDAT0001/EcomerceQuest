package address

import (
	"context"

	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/address/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
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

		err := s.addressUC.CreateAddress(newCtx, &input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Create Address success")
	}
}
