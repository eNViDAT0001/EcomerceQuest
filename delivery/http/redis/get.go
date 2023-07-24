package redis

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/eNViDAT0001/Thesis/Backend/external/cache"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func (s *redisHandler) Get() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		key := cc.Query("key")

		val, err := cache.GetRedis().Get(newCtx, key)
		if err == redis.Nil {
			cc.ResponseError(errors.New("key not found"))
			return
		}
		if err != nil {
			cc.ResponseError(err)
			return
		}

		var value interface{}
		if err = json.Unmarshal([]byte(val), &value); err != nil {
			cc.ResponseError(err)
			return
		}

		result := map[string]interface{}{
			key: value,
		}
		cc.Ok(result)
	}
}
