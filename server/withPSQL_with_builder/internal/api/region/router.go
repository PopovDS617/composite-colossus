package region

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (api *RegionAPI) Router(ctx context.Context) http.Handler {
	r := chi.NewRouter()

	// r.Get("/regions/{id}", api.GetRegionHandler(ctx))
	r.Get("/", api.GetAllRegionsHandler(ctx))
	// r.Post("/regions", api.CreateRegionHandler(ctx))
	// r.Delete("/regions/{id}", api.DeleteRegionHandler(ctx))
	// r.Patch("/regions/{id}", api.UpdateRegionHandler(ctx))

	return r
}
