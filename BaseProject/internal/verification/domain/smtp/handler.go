package smtp

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	CreateResetPassCode() func(*gin.Context)
	VerifyCode() func(*gin.Context)
	CreateEmailFeedback() func(*gin.Context)
}
