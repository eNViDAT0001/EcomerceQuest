package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_viper"
	"github.com/redis/go-redis/v9"
	"time"
)

var v = wrap_viper.GetViper()

func GetRedis() Cache {
	host := v.GetString("REDIS.HOST")
	port := v.GetString("REDIS.PORT")
	db := v.GetInt("REDIS.DB")
	password := v.GetString("REDIS.PASSWORD")

	return NewRedisCache(fmt.Sprintf("%s:%s", host, port), db, password)
}

type redisCache struct {
	db *redis.Client
}

func (r *redisCache) SetDefault(ctx context.Context, key string, value interface{}) error {
	js, err := json.Marshal(value)
	if err != nil {
		return err
	}

	r.db.Set(ctx, key, js, v.GetDuration("REDIS.EXPIRED"))
	return nil
}

func (r *redisCache) Set(ctx context.Context, key string, value interface{}, expired time.Duration) error {
	defer r.close()

	js, err := json.Marshal(value)
	if err != nil {
		return err
	}

	r.db.Set(ctx, key, js, expired)
	return nil
}

func (r redisCache) Get(ctx context.Context, key string) (string, error) {
	defer r.close()
	val, err := r.db.Get(ctx, key).Result()
	return val, err
}
func (r redisCache) close() {
	r.db.Close()
}

func NewRedisCache(host string, db int, password string) Cache {
	return &redisCache{
		db: redis.NewClient(&redis.Options{
			Addr:     host,
			Password: password,
			DB:       db,
		}),
	}
}
