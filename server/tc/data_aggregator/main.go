package main

import (
	"data_aggregator/handlers"
	"data_aggregator/middleware"
	"data_aggregator/pb"
	"data_aggregator/service"
	"data_aggregator/store"
	"data_aggregator/transport"
	"os"

	"fmt"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {

	var (
		httpListenAddress = os.Getenv("HTTP_PORT")
		grpcListenAddress = os.Getenv("GRPC_PORT")
	)

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
	})

	var svc service.Aggregator
	store := store.NewMemoryStore()

	svc = service.NewInvoiceAggregator(store)
	svc = middleware.NewLogMiddleware(svc)
	svc = middleware.NewMetricsMiddleware(svc)

	go makeGRPCTransport(grpcListenAddress, svc)
	makeHTTPTransport(httpListenAddress, svc)

}

func makeHTTPTransport(port string, svc service.Aggregator) {

	aggregateMetricHandler := handlers.NewHTTPMetricHandler("aggregate")
	calculateMetricHandler := handlers.NewHTTPMetricHandler("calculate")

	fmt.Println("http transport running on port", port)
	http.HandleFunc("/aggregator", aggregateMetricHandler.Instrument(handlers.HandleAggregate(svc)))
	http.HandleFunc("/invoice", calculateMetricHandler.Instrument(handlers.HandleGetInvoice(svc)))
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func makeGRPCTransport(listenAddr string, svc service.Aggregator) error {
	fmt.Println("GRPC transport running on port", listenAddr)

	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", listenAddr))
	if err != nil {
		return err
	}
	defer func() {
		fmt.Println("stopping GRPC transport")
		ln.Close()
	}()
	server := grpc.NewServer([]grpc.ServerOption{}...)

	pb.RegisterAggregatorServer(server, transport.NewGRPCAggregatorServer(svc))
	return server.Serve(ln)
}
