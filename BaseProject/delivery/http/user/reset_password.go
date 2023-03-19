package user

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/user/io"
	"gorm.io/gorm"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s userHandler) ResetPassword() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		id, _ := strconv.Atoi(cc.Param("user_id"))

		var input io.ResetPasswordReq
		if err := cc.BindJSON(&input); err != nil {
			cc.ResponseError(err)
			return
		}

		err := s.userUC.ResetPassword(newCtx, uint(id), input.Email)

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.ResponseError(request.NewConflictError("password", "", "password does not match"))
				return
			}
			cc.ResponseError(err)
			return
		}

		cc.Ok("Update password success")
	}
}
