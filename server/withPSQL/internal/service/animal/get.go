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
	var animals []*model.Animal
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		animals, errTx = s.animalRepository.GetAll(ctx)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return animals, nil
}
