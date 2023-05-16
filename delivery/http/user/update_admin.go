package user

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/user/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s userHandler) AdminUpdateUser() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.UpdateUserFullDetailReq
		err := UpdateUser(cc.Context, newCtx, &s, input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Update user success")
	}
}
