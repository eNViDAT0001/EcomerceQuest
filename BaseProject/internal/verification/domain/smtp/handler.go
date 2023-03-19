package smtp

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	CreateResetPassCode() func(*gin.Context)
	VerifyCode() func(*gin.Context)
	GetReportEmail() func(*gin.Context)
	GetFeedbackEmail() func(*gin.Context)
}
