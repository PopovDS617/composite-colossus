package region

import (
	"context"

	"withpsql/internal/model"
)

// func (s *serv) Get(ctx context.Context, id int64) (*model.Animal, error) {
// 	animal, err := s.animalRepository.Get(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return animal, nil
// }

func (s *serv) GetAll(ctx context.Context) ([]*model.Region, error) {
	var regions []*model.Region
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		regions, errTx = s.regionRepository.GetAll(ctx)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return regions, nil
}
