package endpoint

import (
	"context"
	"data_aggregator/service"
	"data_aggregator/types"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

type Set struct {
	AggregateEndoint  endpoint.Endpoint
	CalculateEndpoint endpoint.Endpoint
}

type AggregateRequest struct {
	Value float64 `json:"value"`
	OBUID int     `json:"obu_id"`
	Unix  int64   `json:"timestamp"`
}

type AggregateResponse struct {
	Err error `json:"err"`
}

type CalculateRequest struct {
	OBUID int `json:"obu_id"`
}

type CalculateResponse struct {
	OBUID         int     `json:"obu_id"`
	TotalDistance float64 `json:"total_distance"`
	TotalAmount   float64 `json:"total_amount"`
	Err           error   `json:"err"`
}

func (s Set) Aggregate(ctx context.Context, distance types.Distance) error {
	_, err := s.AggregateEndoint(ctx, AggregateRequest{
		OBUID: distance.OBUID,
		Value: distance.Value,
		Unix:  distance.Unix})

	return err
}

func (s Set) Calculate(ctx context.Context, obuID int) (*types.Invoice, error) {
	res, err := s.AggregateEndoint(ctx, CalculateRequest{
		OBUID: obuID})

	if err != nil {
		return nil, err
	}

	result := res.(CalculateResponse)

	return &types.Invoice{
		OBUID:         result.OBUID,
		TotalDistance: result.TotalAmount,
		TotalAmount:   result.TotalAmount,
	}, nil
}

func New(svc service.Service, logger log.Logger) Set {

	aggregateEndpoint := makeAggregateEndpoint(svc)
	calculateEndpoint := makeCalculateEndpoint(svc)

	return Set{
		AggregateEndoint:  aggregateEndpoint,
		CalculateEndpoint: calculateEndpoint,
	}
}

func makeAggregateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		req := request.(AggregateRequest)

		err = s.Aggregate(ctx, types.Distance{
			OBUID: req.OBUID,
			Value: req.Value,
			Unix:  req.Unix,
		})

		return AggregateResponse{Err: err}, nil
	}
}

func makeCalculateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		req := request.(CalculateRequest)

		inv, err := s.Calculate(ctx, req.OBUID)

		return CalculateResponse{
			OBUID:         req.OBUID,
			TotalDistance: inv.TotalDistance,
			TotalAmount:   inv.TotalAmount,
			Err:           err}, nil
	}
}
