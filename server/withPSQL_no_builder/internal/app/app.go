package app

import (
	"context"

	"withpsql/internal/closer"
	"withpsql/internal/server"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	serviceProvider *serviceProvider
	httpServer      *server.HTTPServer
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

	animal := a.serviceProvider.AnimalAPI(ctx)
	animalRouter := animal.Router(ctx)

	r.Use(middleware.Recoverer)

	r.Mount("/animals", animalRouter)

	port := a.serviceProvider.HTTPConfig().Port()

	server := server.NewHTTPServer(r, port)

	a.httpServer = server

	return nil
}

func (a *App) runHTTPServer() error {
	a.httpServer.PrintStatus()

	closer.Add(a.httpServer.Shutdown)

	return a.httpServer.Run()
}
