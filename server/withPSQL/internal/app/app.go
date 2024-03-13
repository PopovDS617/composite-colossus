package app

import (
	"context"
	"net/http"

	"withpsql/internal/closer"
)

type App struct {
	serviceProvider *serviceProvider
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return http.ListenAndServe("localhost:9000", a.serviceProvider.animalImpl.Router(context.Background()))
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{

		a.initServiceProvider,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}
