package chat

import (
	"github.com/eNViDAT0001/Thesis/Backend/socket/io"
	"github.com/gin-gonic/gin"
)

type HttpHandler interface {
	CreateMessage() func(ctx *gin.Context)
	UpdateMessages() func(ctx *gin.Context)
	SeenMessages() func(ctx *gin.Context)
	ListMessages() func(ctx *gin.Context)
}

type WebSocketHandler interface {
	UnseenMessages() io.EventHandler
	NewMessage() io.EventHandler
	SeenMessage() io.EventHandler
}
