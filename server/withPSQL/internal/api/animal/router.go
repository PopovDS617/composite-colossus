package animal

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (i *Implementation) Router(ctx context.Context) http.Handler {
	r := chi.NewRouter()

	r.Get("/animals/{id}", i.GetAnimalHandler(ctx))
	r.Get("/animals", i.GetAllAnimalsHandler(ctx))
	r.Post("/animals", i.CreateAnimalHandler(ctx))
	r.Delete("/animals/{id}", i.DeleteAnimalHandler(ctx))
	r.Patch("/animals/{id}", i.UpdateAnimalHandler(ctx))

	return r
}
