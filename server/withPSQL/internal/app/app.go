package app

import (
	"context"
	"net/http"

	"withpsql/internal/closer"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	serviceProvider *serviceProvider
	router          *chi.Mux
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
	return a.runHTTPServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initHTTPServer,
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

func (a *App) initHTTPServer(ctx context.Context) error {

	r := chi.NewRouter()

	animal := a.serviceProvider.AnimalImpl(ctx)
	animalRouter := animal.Router(ctx)

	r.Use(middleware.Recoverer)

	r.Mount("/", animalRouter)

	a.router = r

	return nil
}

func (a *App) runHTTPServer() error {
	return http.ListenAndServe("localhost:9000", a.router)
}
