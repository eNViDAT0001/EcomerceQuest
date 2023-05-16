package smtp

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/verification/smtp/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s *smtpHandler) VerifyCode() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.CodeVerificationForm
		if err := cc.BindJSON(&input); err != nil {
			cc.ResponseError(err)
			return
		}

		token, err := s.jwtUC.VerifySmtpToken(newCtx, input.UserID, input.Token, input.Code)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok(token.Valid)
	}
}
