package jwt

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (s *jwtHandler) VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		newCtx := context.Background()

		cancel := func() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid access_token"})
		}

		bearToken := c.Request.Header.Get("Authorization")
		strArr := strings.Split(bearToken, " ")
		tokenString := ""
		if len(strArr) == 2 {
			tokenString = strArr[1]
		}

		_, err := s.jwtUC.VerifyToken(newCtx, tokenString)
		if err != nil {
			cancel()
			return
		}
	}
}
