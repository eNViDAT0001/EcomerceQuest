package product_quantities

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eNViDAT0001/Thesis/Backend/external/cache"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
)

var couponInstance *couponStorage

type couponStorage struct {
	mu sync.Mutex
}

func GetCouponStore() *couponStorage {
	if instance != nil {
		return couponInstance
	}

	return &couponStorage{}
}

func (fRedis *couponStorage) Minus(ctx context.Context, id uint, quantity int) (ok bool, err error) {
	fRedis.mu.Lock()
	defer fRedis.mu.Unlock()

	key := fmt.Sprintf("store_%d", id)
	value, err := cache.GetRedis().Get(ctx, key)
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		log.Println(err)
		return false, err
	}

	var savedQuantity int
	err = json.Unmarshal([]byte(value), &savedQuantity)

	if err != nil {
		return false, err
	}

	if savedQuantity < quantity {
		return false, fmt.Errorf("not enough quantity")
	}
	cache.GetRedis().SetDefault(ctx, fmt.Sprintf("store_%d", id), savedQuantity-quantity)
	return true, nil
}
func (fRedis *couponStorage) Add(ctx context.Context, id uint, quantity int) {
	fRedis.mu.Lock()
	defer fRedis.mu.Unlock()

	value, err := cache.GetRedis().Get(ctx, fmt.Sprintf("store_%d", id))
	if err != nil && err != redis.Nil {
		log.Println(err)
		return
	}
	var savedQuantity int
	err = json.Unmarshal([]byte(value), &savedQuantity)
	err = cache.GetRedis().SetDefault(ctx, fmt.Sprintf("store_%d", id), savedQuantity+quantity)
	if err != nil {
		log.Println(err)
	}
}

func (fRedis *couponStorage) Reduce(ctx context.Context, store map[uint]int) (okela bool, invalidKey uint) {
	storage := map[uint]int{}
	for k, v := range store {
		ok, err := fRedis.Minus(ctx, k, v)
		if err != nil {
			fRedis.Restore(ctx, storage)
			return false, 0
		}

		if !ok {
			fRedis.Restore(ctx, storage)
			return false, k
		}
		storage[k] = v
	}

	return true, 0
}

func (fRedis *couponStorage) Restore(ctx context.Context, store map[uint]int) {
	for k, v := range store {
		fRedis.Add(ctx, k, v)
	}
}
