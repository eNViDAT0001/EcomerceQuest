package user

import (
	"context"

	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/user/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
)

func (s userHandler) DeleteUserByIDs() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.DeleteUserReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		err := s.userUC.DeleteUserByIDs(newCtx, input.IDs)

		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Delete User success")
	}
}
