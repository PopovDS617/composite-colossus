package animal

import (
	"context"

	"withpsql/internal/model"
)

func (s *serv) Create(ctx context.Context, animal *model.Animal) (int64, error) {
	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.animalRepository.Create(ctx, animal)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.animalRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
