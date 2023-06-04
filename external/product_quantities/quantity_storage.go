package product_quantities

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eNViDAT0001/Thesis/Backend/external/cache"
	"log"
	"sync"
)

var storage *productStorage

type productStorage struct {
	mu      sync.Mutex
	storage map[uint]int
}

func GetQuantityStore() *productStorage {
	if storage != nil {
		return storage
	}

	store := map[uint]int{}
	return &productStorage{storage: store}
}

func (fRedis *productStorage) Minus(id uint, quantity int) (ok bool, err error) {
	fRedis.mu.Lock()
	defer fRedis.mu.Unlock()

	value, err := cache.GetRedis().Get(context.Background(), fmt.Sprintf("store_%d", id))
	if err != nil {
		log.Println(err)
		return false, err
	}

	var savedQuantity int
	err = json.Unmarshal([]byte(value), &savedQuantity)

	if err != nil {
		return false, err
	}
	if savedQuantity == 0 {
		return false, nil
	}

	if savedQuantity < quantity {
		return false, fmt.Errorf("not enough quantity")
	}
	cache.GetRedis().SetDefault(context.Background(), fmt.Sprintf("store_%d", id), savedQuantity-quantity)
	return true, nil
}
func (fRedis *productStorage) Add(id uint, quantity int) {
	fRedis.mu.Lock()
	defer fRedis.mu.Unlock()

	value, err := cache.GetRedis().Get(context.Background(), fmt.Sprintf("store_%d", id))
	if err != nil {
		if err != nil {
			log.Println(err)
		}
		return
	}
	var savedQuantity int
	err = json.Unmarshal([]byte(value), &savedQuantity)
	err = cache.GetRedis().SetDefault(context.Background(), fmt.Sprintf("store_%d", id), savedQuantity+quantity)
	if err != nil {
		log.Println(err)
	}
}

func (fRedis *productStorage) Reduce(store map[uint]int) (okela bool, invalidKey uint) {

	storage := map[uint]int{}
	for k, v := range store {
		ok, err := fRedis.Minus(k, v)
		if err != nil {
			fRedis.Restore(storage)
			return false, 0
		}

		if !ok {
			fRedis.Restore(storage)
			return false, k
		}
		storage[k] = v
	}

	return true, 0
}

func (fRedis *productStorage) Restore(store map[uint]int) {
	for k, v := range store {
		fRedis.Add(k, v)
	}
}
