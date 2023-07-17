package redis

import "github.com/gin-gonic/gin"

type DummyRedisHandler interface {
	Get() func(ctx *gin.Context)
	Set() func(ctx *gin.Context)
}

type redisHandler struct {
}

func NewRedisHandler() DummyRedisHandler {
	return &redisHandler{}
}
