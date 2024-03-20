package region

import (
	"withpsql/internal/client/db"
	"withpsql/internal/repository"
	"withpsql/internal/service"
)

type serv struct {
	regionRepository repository.RegionRepository
	txManager        db.TxManager
}

func NewService(
	regionRepository repository.RegionRepository,
	txManager db.TxManager,
) service.RegionService {
	return &serv{
		regionRepository: regionRepository,
		txManager:        txManager,
	}
}

func NewMockService(deps ...interface{}) service.RegionService {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.RegionRepository:
			srv.regionRepository = s
		}
	}

	return &srv
}
