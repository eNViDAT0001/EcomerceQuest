package notification

import (
	"github.com/eNViDAT0001/Thesis/Backend/socket/io"
	"github.com/gin-gonic/gin"
)

type HttpHandler interface {
	ListNotifications() func(ctx *gin.Context)
	SeenNotification() func(ctx *gin.Context)
}
type WebSocketHandler interface {
	UnseenNotification() io.EventHandler
	NewNotification() io.EventHandler
}
