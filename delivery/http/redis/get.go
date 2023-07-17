package redis

import (
	"context"
	"encoding/json"
	"github.com/eNViDAT0001/Thesis/Backend/external/cache"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s *redisHandler) Get() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		key := cc.Query("key")

		val, err := cache.GetRedis().Get(newCtx, key)
		if err != nil {
			cc.ResponseError(err)
		}

		var value string
		if err = json.Unmarshal([]byte(val), &value); err != nil {
			cc.ResponseError(err)
		}

		result := map[string]interface{}{
			key: value,
		}
		cc.Ok(result)
	}
}
