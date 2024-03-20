package repository

import (
	"context"

	model "withpsql/internal/model"
)

type AnimalRepository interface {
	Create(ctx context.Context, animal *model.Animal) (int64, error)
	Get(ctx context.Context, id int64) (*model.Animal, error)
}
