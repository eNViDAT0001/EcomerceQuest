package cache

import (
	"context"
	"time"
)

type Cache interface {
	Set(ctx context.Context, key string, value interface{}, expired time.Duration) error
	SetDefault(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (string, error)
}
