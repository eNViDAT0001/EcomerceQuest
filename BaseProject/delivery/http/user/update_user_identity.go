package user

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/user/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s userHandler) UpdateUserIdentity() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.UpdateUserIdentityReq

		id, _ := strconv.Atoi(cc.Param("user_id"))
		_, err := s.userUC.ComparePassword(newCtx, uint32(id), input.Password)

		err = UpdateUser(cc.Context, newCtx, &s, input)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		cc.Ok("Update user success")
	}
}
