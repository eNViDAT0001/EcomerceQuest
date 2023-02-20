package jwt

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	VerifyToken() gin.HandlerFunc
	VerifyAdminToken() gin.HandlerFunc
	VerifyUserToken() gin.HandlerFunc
	VerifyProductWithUserToken() gin.HandlerFunc
	VerifyProviderWithUserToken() gin.HandlerFunc
	RefreshToken() gin.HandlerFunc
}
