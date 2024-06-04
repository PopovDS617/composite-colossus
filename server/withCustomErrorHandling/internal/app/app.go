package app

import (
	"net/http"
	repo "withcustomerrorhandling/internal/repository"
	carRepo "withcustomerrorhandling/internal/repository/car"
	"withcustomerrorhandling/internal/server"
	svc "withcustomerrorhandling/internal/service"
	carSvc "withcustomerrorhandling/internal/service/car"
	carUsecase "withcustomerrorhandling/internal/usecase/car"
	"withcustomerrorhandling/pkg/handlers"
)

type App struct {
	router        http.Handler
	httpServer    server.Server
	carService    svc.CarService
	carRepository repo.CarRepository
	carUsecase    *carUsecase.CarUsecase
}

func (a *App) initRouter() http.Handler {
	if a.router == nil {

		mux := http.NewServeMux()

		mux.Handle("/car/{id}/", handlers.CustomHandler(a.carUsecase.Get))

		a.router = mux
	}
	return a.router
}

func (a *App) initServer() server.Server {
	if a.httpServer == nil {
		a.httpServer = server.NewHTTPServer(a.router)
	}

	return a.httpServer
}

func (a *App) initCarService() svc.CarService {
	if a.carService == nil {
		a.carService = carSvc.NewCarService(a.carRepository)
	}
	return a.carService
}

func (a *App) initCarRepository() repo.CarRepository {
	if a.carRepository == nil {
		a.carRepository = carRepo.NewCarRepository()
	}

	return a.carRepository
}

func (a *App) initCarUsecase() carUsecase.CarUsecase {
	if a.carUsecase == nil {
		a.carUsecase = carUsecase.NewCarUsecase(a.carService)
	}
	return *a.carUsecase
}

func NewApp() *App {
	app := &App{}

	app.initCarRepository()
	app.initCarService()
	app.initCarUsecase()
	app.initRouter()
	app.initServer()

	return app
}

func (a *App) Start() error {
	return a.httpServer.Run()
}
