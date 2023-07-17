package redis

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/cache"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s *redisHandler) Set() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		key := cc.Query("key")
		value := cc.Query("value")

		err := cache.GetRedis().SetDefault(newCtx, key, value)
		if err != nil {
			cc.ResponseError(err)
		}

		result := map[string]interface{}{
			key: value,
		}
		cc.Ok(result)
	}
}
