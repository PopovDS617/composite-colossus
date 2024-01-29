package service

import (
	"data_aggregator/types"
	"fmt"
)

type Aggregator interface {
	AggregateDistance(types.Distance) error
}

type Storer interface {
	Put(types.Distance) error
}

type InvoiceAggregator struct {
	store Storer
}

func NewInvoiceAggregator(store Storer) *InvoiceAggregator {
	return &InvoiceAggregator{
		store: store,
	}
}

func (ia *InvoiceAggregator) AggregateDistance(data types.Distance) error {
	fmt.Println("inserting distance in the storage")
	return ia.store.Put(data)
}
