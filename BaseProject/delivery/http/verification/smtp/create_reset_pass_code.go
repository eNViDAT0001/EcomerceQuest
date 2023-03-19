package smtp

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"strconv"
)

func (s *smtpHandler) CreateResetPassCode() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		id, _ := strconv.Atoi(cc.Param("user_id"))
		email := cc.Param("user_id")

		token, code, err := s.jwtUC.GenerateSmtpCode(newCtx, uint(id))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		err = s.smtpUC.SendEmail(newCtx, email, code)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok(token)
	}
}
