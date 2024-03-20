package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
)

type HTTPServer struct {
	server *http.Server
}

func NewHTTPServer(router *chi.Mux, port string) *HTTPServer {
	return &HTTPServer{
		server: &http.Server{
			Handler: router,
			Addr:    fmt.Sprintf(":%s", port),
		},
	}
}

func (s *HTTPServer) Run() error {
	return s.server.ListenAndServe()
}

func (s *HTTPServer) PrintStatus() {
	fmt.Println("server listening on port", s.server.Addr)

}
func (s *HTTPServer) Shutdown() error {

	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	defer stop()

	fmt.Println("shutting down the server")

	return s.server.Shutdown(ctx)

}
