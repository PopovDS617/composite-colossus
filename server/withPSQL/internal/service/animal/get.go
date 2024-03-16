package animal

import (
	"context"

	"withpsql/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.Animal, error) {
	animal, err := s.animalRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return animal, nil
}

func (s *serv) GetAll(ctx context.Context) ([]*model.Animal, error) {
	animals, err := s.animalRepository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return animals, nil
}
