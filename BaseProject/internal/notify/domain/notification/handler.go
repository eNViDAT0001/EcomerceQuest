package notification

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	ListNotifications() func(ctx *gin.Context)
}
