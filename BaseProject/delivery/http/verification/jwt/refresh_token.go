package jwt

import (
	"context"

	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s *jwtHandler) RefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var tokenReq RefreshTokenReq
		if err := cc.BindJSON(&tokenReq); err != nil {
			cc.ResponseError(err)
			return
		}

		token, err := s.jwtUC.RefreshToken(newCtx, tokenReq.RefreshToken)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok(token)
	}
}

type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}
