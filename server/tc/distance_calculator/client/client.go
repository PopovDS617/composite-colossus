package client

import (
	"context"
	"dist_calc/pb"
)

type Client interface {
	Aggregate(context.Context, *pb.AggregateRequest) error
}
