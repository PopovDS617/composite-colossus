package animal

import (
	"context"
	"net/http"
)

func (i *Implementation) Router(ctx context.Context) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /animal/{id}", i.GetAnimalHandler(ctx))
	mux.HandleFunc("POST /animal", i.CreateAnimalHandler(ctx))

	return mux
}
