package transport

import (
	"data_aggregator/pb"
	"data_aggregator/service"
	"data_aggregator/types"
)

type GRPCAggregatorServer struct {
	pb.UnimplementedAggregatorServer
	svc service.Aggregator
}

func NewGRPCAggregatorServer(svc service.Aggregator) *GRPCAggregatorServer {
	return &GRPCAggregatorServer{
		svc: svc,
	}
}

func (s *GRPCAggregatorServer) AggregateDistance(req pb.AggregateRequest) error {
	distance := types.Distance{
		OBUID: int(req.OBUID),
		Value: req.Value,
		Unix:  req.Unix,
	}

	return s.svc.AggregateDistance(distance)
}
