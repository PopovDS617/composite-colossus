package transport

import (
	"context"
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

func (s *GRPCAggregatorServer) AggregateDistance(ctx context.Context, req *pb.AggregateRequest) (*pb.None, error) {
	distance := types.Distance{
		OBUID: int(req.OBUID),
		Value: req.Value,
		Unix:  req.Unix,
	}

	return &pb.None{}, s.svc.AggregateDistance(distance)
}
