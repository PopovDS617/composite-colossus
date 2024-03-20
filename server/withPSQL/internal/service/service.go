package service

import (
	"context"

	"withpsql/internal/model"
)

type AnimalService interface {
	Create(ctx context.Context, animal *model.Animal) (int64, error)
	Get(ctx context.Context, id int64) (*model.Animal, error)
}
