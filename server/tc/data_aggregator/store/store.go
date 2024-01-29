package store

import "data_aggregator/types"

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
