package animal

import (
	"context"

	"withpsql/internal/model"
)

func (s *serv) Update(ctx context.Context, animal *model.Animal) (*model.Animal, error) {

	var updatedAnimal *model.Animal
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.animalRepository.Update(ctx, animal)
		if errTx != nil {
			return errTx
		}

		updatedAnimal, errTx = s.animalRepository.Get(ctx, animal.ID)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return updatedAnimal, nil
}
