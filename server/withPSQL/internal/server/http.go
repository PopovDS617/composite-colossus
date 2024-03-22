package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HTTPServer struct {
	Router *chi.Mux
	Port   string
}

func NewHTTPServer(router *chi.Mux, port string) *HTTPServer {
	return &HTTPServer{
		Router: router,
		Port:   port,
	}
}

func (s *HTTPServer) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%s", s.Port), s.Router)
}
