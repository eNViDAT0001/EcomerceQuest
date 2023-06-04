package user

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/user/io"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/verification/smtp"

	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s userHandler) ResetPassword() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		userEmail := cc.Param("email")
		user, err := s.userUC.GetUserByEmail(newCtx, userEmail)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		var input io.ResetPasswordReq
		if err := cc.BindJSON(&input); err != nil {
			cc.ResponseError(err)
			return
		}

		token, err := s.jwtUC.VerifySmtpToken(newCtx, user.ID, input.Token, input.Code)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		if !token.Valid {
			cc.ResponseError(request.NewUnauthorizedError("Code", input.Code, "Code is not valid"))
			return
		}

		ok := smtp.UseToken(newCtx, input.Token)
		if !ok {
			cc.ResponseError(request.NewUnauthorizedError("Code", input.Code, "Code not found"))
			return
		}

		err = s.userUC.ResetPassword(newCtx, user.ID, input.NewPassword)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Update password success")
	}
}
