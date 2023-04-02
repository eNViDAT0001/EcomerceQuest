package user

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/user/io"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)
var invalidTokens = make(map[string]bool)
func (s userHandler) ResetPassword() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		userID, _ := strconv.Atoi(cc.Param("user_id"))

		var input io.ResetPasswordReq
		if err := cc.BindJSON(&input); err != nil {
			cc.ResponseError(err)
			return
		}

		token, err := s.jwtUC.VerifySmtpToken(newCtx, uint(userID), input.Token, input.Code)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		if !token.Valid {
			cc.ResponseError(request.NewUnauthorizedError("Code", input.Code, "Code is not valid"))
			return
		}

		if _, ok := invalidTokens[input.Token]; ok {
			cc.ResponseError(request.NewUnauthorizedError("Code", input.Code, "Code is not valid"))
		}
		invalidTokens[input.Token] = true

		err = s.userUC.ResetPassword(newCtx, uint(userID), input.NewPassword)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Update password success")
	}
}
