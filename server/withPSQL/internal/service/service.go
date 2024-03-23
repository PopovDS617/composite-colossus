package service

import (
	"context"

	"withpsql/internal/model"
)

type AnimalService interface {
	Create(ctx context.Context, animal *model.Animal) (*model.Animal, error)
	Get(ctx context.Context, id int64) (*model.Animal, error)
	GetAll(ctx context.Context) ([]*model.Animal, error)
}
