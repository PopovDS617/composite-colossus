package animal

import (
	"context"

	"withpsql/internal/model"
)

func (s *serv) Create(ctx context.Context, animal *model.Animal) (*model.Animal, error) {
	var id int64
	var insertedAnimal *model.Animal
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.animalRepository.Create(ctx, animal)
		if errTx != nil {
			return errTx
		}

		insertedAnimal, errTx = s.animalRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return insertedAnimal, nil
}
