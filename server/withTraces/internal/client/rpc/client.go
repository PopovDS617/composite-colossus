package rpc

import (
	"context"

	"withtraces/internal/model"
)

type OtherServiceClient interface {
	Get(ctx context.Context, id int64) (*model.Note, error)
}
