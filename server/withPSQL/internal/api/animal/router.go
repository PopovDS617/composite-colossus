package animal

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (api *AnimalAPI) Router(ctx context.Context) http.Handler {
	r := chi.NewRouter()

	r.Get("/{id}", api.GetAnimalHandler(ctx))
	r.Get("/", api.GetAllAnimalsHandler(ctx))
	r.Post("/", api.CreateAnimalHandler(ctx))
	r.Delete("/{id}", api.DeleteAnimalHandler(ctx))
	r.Patch("/{id}", api.UpdateAnimalHandler(ctx))

	return r
}
