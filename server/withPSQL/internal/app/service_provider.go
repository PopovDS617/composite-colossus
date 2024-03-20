package app

import (
	"context"
	"log"

	animalAPI "withpsql/internal/api/animal"
	regionAPI "withpsql/internal/api/region"
	"withpsql/internal/repository"

	"withpsql/internal/client/db"
	"withpsql/internal/client/db/pg"
	"withpsql/internal/client/db/transaction"
	"withpsql/internal/closer"
	"withpsql/internal/config"
	animalRepo "withpsql/internal/repository/animal"
	regionRepo "withpsql/internal/repository/region"

	svc "withpsql/internal/service"
	animalSvc "withpsql/internal/service/animal"
	regionSvc "withpsql/internal/service/region"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	httpConfig config.HTTPConfig

	dbClient         db.Client
	txManager        db.TxManager
	animalRepository repository.AnimalRepository
	animalService    svc.AnimalService
	animalAPI        *animalAPI.AnimalAPI
	regionRepository repository.RegionRepository
	regionService    svc.RegionService
	regionAPI        *regionAPI.RegionAPI
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) PGConfig() config.PGConfig {

	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {

	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) AnimalRepository(ctx context.Context) repository.AnimalRepository {
	if s.animalRepository == nil {
		s.animalRepository = animalRepo.NewRepository(s.DBClient(ctx))
	}

	return s.animalRepository
}

func (s *serviceProvider) AnimalService(ctx context.Context) svc.AnimalService {
	if s.animalService == nil {
		s.animalService = animalSvc.NewService(
			s.AnimalRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.animalService
}

func (s *serviceProvider) AnimalAPI(ctx context.Context) *animalAPI.AnimalAPI {
	if s.animalAPI == nil {
		s.animalAPI = animalAPI.NewAnimalAPI(s.AnimalService(ctx))

	}

	return s.animalAPI
}
func (s *serviceProvider) RegionRepository(ctx context.Context) repository.RegionRepository {
	if s.regionRepository == nil {
		s.regionRepository = regionRepo.NewRepository(s.DBClient(ctx))
	}

	return s.regionRepository
}

func (s *serviceProvider) RegionService(ctx context.Context) svc.RegionService {
	if s.regionService == nil {
		s.regionService = regionSvc.NewService(
			s.RegionRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.regionService
}

func (s *serviceProvider) RegionAPI(ctx context.Context) *regionAPI.RegionAPI {
	if s.regionAPI == nil {
		s.regionAPI = regionAPI.NewRegionAPI(s.RegionService(ctx))

	}

	return s.regionAPI
}
