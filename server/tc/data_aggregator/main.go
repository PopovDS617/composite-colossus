package main

import (
	"data_aggregator/service"
	endpoint "data_aggregator/service/aggregate-endpoint"
	"data_aggregator/store"
	"data_aggregator/transport"
	"fmt"
	"os"

	"net"
	"net/http"

	"github.com/go-kit/log"
	"github.com/sirupsen/logrus"
)

func main() {

	var (
		httpListenAddress = os.Getenv("HTTP_PORT")
		// grpcListenAddress = os.Getenv("GRPC_PORT")
	)

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
	})

	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	store := store.NewMemoryStore()

	service := service.NewAggregatorService(store)
	endpoints := endpoint.New(service, logger)
	httpHandler := transport.NewHTTPHandler(endpoints, logger)

	httpListener, err := net.Listen("tcp", fmt.Sprintf(":%s", httpListenAddress))
	if err != nil {
		logger.Log("transport", "HTTP", "during", "Listen", "err", err)
		os.Exit(1)
	}
	logger.Log("transport", "HTTP", "addr", fmt.Sprintf(":%s", httpListenAddress))
	err = http.Serve(httpListener, httpHandler)
	if err != nil {
		panic(err)
	}
}
