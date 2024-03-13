package app

import (
	"context"
	"log"

	animalApi "withpsql/internal/api/animal"
	"withpsql/internal/repository"

	"withpsql/internal/client/db"
	"withpsql/internal/client/db/pg"
	"withpsql/internal/client/db/transaction"
	"withpsql/internal/closer"
	"withpsql/internal/config"
	animalRepo "withpsql/internal/repository/animal"
	svc "withpsql/internal/service"
	animalSvc "withpsql/internal/service/animal"
)

type serviceProvider struct {
	pgConfig         config.PGConfig
	dbClient         db.Client
	txManager        db.TxManager
	animalRepository repository.AnimalRepository
	animalService    svc.AnimalService
	animalImpl       *animalApi.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
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

func (s *serviceProvider) AnimalImpl(ctx context.Context) *animalApi.Implementation {
	if s.animalImpl == nil {
		s.animalImpl = animalApi.NewImplementation(s.AnimalService(ctx))
		s.animalImpl.Router(ctx)
	}

	return s.animalImpl
}
