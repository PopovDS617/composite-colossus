package service

import (
	"context"

	"data_aggregator/types"
)

type LoggingMiddleware struct {
	next Service
}

type Middleware func(Service) Service

func NewLoggingMiddleware() Middleware {
	return func(next Service) Service {
		return &LoggingMiddleware{
			next: next,
		}
	}
}

func (lmw *LoggingMiddleware) Aggregate(_ context.Context, distance types.Distance) error {
	return nil
}

func (lmw *LoggingMiddleware) Calculate(_ context.Context, obuID int) (*types.Invoice, error) {
	return nil, nil
}

type InstrumentationMiddleware struct {
	next Service
}

func NewInstrumentationMiddleware() Middleware {
	return func(next Service) Service {
		return &InstrumentationMiddleware{
			next: next,
		}
	}
}

func (imw *InstrumentationMiddleware) Aggregate(_ context.Context, distance types.Distance) error {
	return nil
}

func (imw *InstrumentationMiddleware) Calculate(_ context.Context, obuID int) (*types.Invoice, error) {
	return nil, nil
}
