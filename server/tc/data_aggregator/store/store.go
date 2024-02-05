package store

import (
	"data_aggregator/types"
	"fmt"
)

type Storer interface {
	Put(types.Distance) error
	Get(int) (float64, error)
}

type MemoryStore struct {
	data map[int]float64
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{data: make(map[int]float64)}
}

func (ms *MemoryStore) Put(d types.Distance) error {
	ms.data[d.OBUID] += d.Value
	return nil
}
func (ms *MemoryStore) Get(id int) (float64, error) {
	dist, ok := ms.data[id]

	if !ok {
		return 0.0, fmt.Errorf("data not found")
	}

	return dist, nil

}
