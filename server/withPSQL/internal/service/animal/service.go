package animal

import (
	"withpsql/internal/client/db"
	"withpsql/internal/repository"
	"withpsql/internal/service"
)

type serv struct {
	animalRepository repository.AnimalRepository
	txManager        db.TxManager
}

func NewService(
	animalRepository repository.AnimalRepository,
	txManager db.TxManager,
) service.AnimalService {
	return &serv{
		animalRepository: animalRepository,
		txManager:        txManager,
	}
}

func NewMockService(deps ...interface{}) service.AnimalService {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.AnimalRepository:
			srv.animalRepository = s
		}
	}

	return &srv
}
