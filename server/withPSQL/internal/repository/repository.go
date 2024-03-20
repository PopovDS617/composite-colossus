package repository

import (
	"context"

	model "withpsql/internal/model"
)

type AnimalRepository interface {
	Create(ctx context.Context, animal *model.Animal) (int64, error)
	Update(ctx context.Context, animal *model.Animal) error
	Get(ctx context.Context, id int64) (*model.Animal, error)
	GetAll(ctx context.Context) ([]*model.Animal, error)
	Delete(ctx context.Context, id int64) error
}
type RegionRepository interface {
	// Create(ctx context.Context, animal *model.Animal) (int64, error)
	// Update(ctx context.Context, animal *model.Animal) error
	// Get(ctx context.Context, id int64) (*model.Animal, error)
	GetAll(ctx context.Context) ([]*model.Region, error)
	// Delete(ctx context.Context, id int64) error
}
