package notification

import (
	"github.com/gin-gonic/gin"
)

type HttpHandler interface {
	ListNotifications() func(ctx *gin.Context)
}

//type WebSocketHandler interface {
//	OnConnect(event websocket.Event, client *websocket.Client) error
//	OnNewNotification(event websocket.Event, client *websocket.Client) error
//	OnSeenNotification(event websocket.Event, client *websocket.Client) error
//}
