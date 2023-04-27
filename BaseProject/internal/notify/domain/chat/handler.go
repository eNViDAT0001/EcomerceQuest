package chat

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	CreateMessage() func(ctx *gin.Context)
	UpdateMessages() func(ctx *gin.Context)
	SeenMessages() func(ctx *gin.Context)
	ListMessages() func(ctx *gin.Context)
}
