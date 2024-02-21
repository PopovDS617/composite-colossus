package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/brianvoe/gofakeit"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"

	"withRateLimiter/internal/interceptor"
	"withRateLimiter/internal/metric"
	"withRateLimiter/internal/withRateLimiter"
	desc "withRateLimiter/pkg/note_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedNoteV1Server
}

// Get ...
func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Note id: %d", req.GetId())

	return &desc.GetResponse{
		Note: &desc.Note{
			Id: req.GetId(),
			Info: &desc.NoteInfo{
				Title:    gofakeit.BeerName(),
				Context:  gofakeit.IPv4Address(),
				Author:   gofakeit.Name(),
				IsPublic: gofakeit.Bool(),
			},
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

func main() {
	ctx := context.Background()
	err := metric.Init(ctx)
	if err != nil {
		log.Fatalf("failed to init metrics: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rateLimiter := withRateLimiter.NewTokenBucketLimiter(ctx, 10, time.Second)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				interceptor.NewRateLimiterInterceptor(rateLimiter).Unary,
				interceptor.MetricsInterceptor,
			),
		),
	)
	reflection.Register(s)
	desc.RegisterNoteV1Server(s, &server{})

	go func() {
		err = runPrometheus()
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func runPrometheus() error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	prometheusServer := &http.Server{
		Addr:    "localhost:2112",
		Handler: mux,
	}

	log.Printf("Prometheus server is running on %s", "localhost:2112")

	err := prometheusServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
