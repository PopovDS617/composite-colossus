package client

import (
	"dist_calc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	Endpoint string
	pb.AggregatorClient
}

func NewGRPCClient(endpoint string) (*GRPCClient, error) {
	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	c := pb.NewAggregatorClient(conn)

	if err != nil {
		return nil, err
	}

	return &GRPCClient{Endpoint: endpoint, AggregatorClient: c}, nil
}
