package smtp

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	CreateResetPassCode() func(*gin.Context)
	VerifyCode() func(*gin.Context)
	CreateEmailFeedback() func(*gin.Context)
	List() func(*gin.Context)
	SendEmail() func(*gin.Context)

	ReSendUnSendEmail() error
	ReSendFailedEmail() func(*gin.Context)
}
