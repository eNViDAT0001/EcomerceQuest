package jwt

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/user/entities"

	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (s *jwtHandler) VerifyProviderWithUserToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		cancel := func() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid access_token"})
		}
		permissionCancel := func() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "out of permission"})
		}

		bearToken := c.Request.Header.Get("Authorization")
		strArr := strings.Split(bearToken, " ")
		tokenString := ""
		if len(strArr) == 2 {
			tokenString = strArr[1]
		}

		token, err := s.jwtUC.VerifyToken(newCtx, tokenString)
		if err != nil {
			cancel()
			return
		}
		claims, _ := token.Claims.(jwt.MapClaims)

		userID, _ := strconv.Atoi(claims["user_id"].(string))
		user, err := s.userUC.GetUserDetailByID(newCtx, uint(userID))

		if *user.Type == entities.Admin {
			return
		}

		providerID, err := strconv.Atoi(cc.Param("provider_id"))
		if err != nil {
			cancel()
			return
		}
		provider, err := s.providerUC.GetProviderByID(newCtx, uint(providerID))
		if err != nil {
			cancel()
			return
		}

		id, err := strconv.Atoi(cc.Param("user_id"))
		if err != nil {
			cancel()
			return
		}
		if uint(id) != user.ID || uint(id) != provider.UserID {
			permissionCancel()
			return
		}
	}
}
