package fake_redis

import (
	"fmt"
	"sync"
)

var storage *fakeRedis

type fakeRedis struct {
	mu      sync.Mutex
	storage map[uint]int
}

func GetQuantityStore() *fakeRedis {
	if storage != nil {
		return storage
	}

	store := map[uint]int{}
	return &fakeRedis{storage: store}
}

func (fRedis *fakeRedis) GetQuantityDict() map[uint]int {
	return fRedis.storage
}

func (fRedis *fakeRedis) Minus(id uint, quantity int) (ok bool, err error) {
	fRedis.mu.Lock()
	defer fRedis.mu.Unlock()

	val, ok := fRedis.storage[id]
	if !ok {
		fRedis.storage[id] = 0
		return false, nil
	}

	if val < quantity {
		return false, fmt.Errorf("not enough quantity")
	}

	fRedis.storage[id] = val - quantity
	return true, nil
}
func (fRedis *fakeRedis) Add(id uint, quantity int) {
	fRedis.mu.Lock()
	defer fRedis.mu.Unlock()

	val, ok := fRedis.storage[id]
	if ok {
		fRedis.storage[id] = val + quantity
		return
	}
	fRedis.storage[id] = quantity
}

func (fRedis *fakeRedis) Reduce(store map[uint]int) (okela bool, invalidKey uint) {

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

func (fRedis *fakeRedis) Restore(store map[uint]int) {
	for k, v := range store {
		fRedis.Add(k, v)
	}
}
