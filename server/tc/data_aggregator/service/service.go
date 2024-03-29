package service

import (
	"context"

	"data_aggregator/store"
	"data_aggregator/types"

	"github.com/go-kit/log"
)

const price = 1.216

type Service interface {
	Aggregate(context.Context, types.Distance) error
	Calculate(context.Context, int) (*types.Invoice, error)
}

type BasicService struct {
	store store.Storer
}

func newBasicService(store store.Storer) Service {
	return &BasicService{
		store}
}

func (svc *BasicService) Aggregate(_ context.Context, distance types.Distance) error {

	return svc.store.Put(distance)
}
func (svc *BasicService) Calculate(_ context.Context, obuID int) (*types.Invoice, error) {

	distance, err := svc.store.Get(obuID)

	if err != nil {
		return nil, err
	}

	inv := &types.Invoice{
		OBUID:         obuID,
		TotalDistance: distance,
		TotalAmount:   distance * price,
	}

	return inv, nil
}

func New(store store.Storer, logger log.Logger) Service {
	var svc Service
	{
		svc = newBasicService(store)
		svc = newLoggingMiddleware(logger)(svc)
		svc = newInstrumentationMiddleware()(svc)
	}
	return svc
}
