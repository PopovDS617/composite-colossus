package service

import (
	"data_aggregator/types"
)

const price = 1.216

type Aggregator interface {
	AggregateDistance(types.Distance) error
	GenerateInvoice(int) (*types.Invoice, error)
}

type Storer interface {
	Put(types.Distance) error
	Get(int) (float64, error)
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

	return ia.store.Put(data)
}

func (ia *InvoiceAggregator) GenerateInvoice(id int) (*types.Invoice, error) {

	distance, err := ia.store.Get(id)

	if err != nil {
		return nil, err
	}

	inv := &types.Invoice{
		OBUID:         id,
		TotalDistance: distance,
		TotalAmount:   distance * price,
	}
	return inv, nil
}
