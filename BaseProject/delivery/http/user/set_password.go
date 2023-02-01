package user

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/user/io"
	"gorm.io/gorm"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s userHandler) SetPassword() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		id, _ := strconv.Atoi(cc.Param("user_id"))

		var password io.PasswordRequest
		if err := cc.BindJSON(&password); err != nil {
			cc.ResponseError(err)
			return
		}

		err := s.userUC.SetPassword(newCtx, uint32(id), password.Password, password.NewPassword)

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
