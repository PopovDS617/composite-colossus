package main

import (
	"data_aggregator/handlers"
	"data_aggregator/middleware"
	"data_aggregator/pb"
	"data_aggregator/service"
	"data_aggregator/store"
	"data_aggregator/transport"

	"fmt"
	"net"
	"net/http"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
	})

	var svc service.Aggregator

	httpListenAddress := ":9000"
	grpcListenAddress := ":9001"

	store := store.NewMemoryStore()

	svc = service.NewInvoiceAggregator(store)
	svc = middleware.NewLogMiddleware(svc)

	go makeGRPCTransport(grpcListenAddress, svc)
	makeHTTPTransport(httpListenAddress, svc)

}

func makeHTTPTransport(port string, svc service.Aggregator) {
	fmt.Println("http transport running on port", port)
	http.HandleFunc("/aggregator", handlers.HandleAggregate(svc))
	http.HandleFunc("/invoice", handlers.HandleGetInvoice(svc))
	http.ListenAndServe(port, nil)
}

func makeGRPCTransport(listenAddr string, svc service.Aggregator) error {
	fmt.Println("GRPC transport running on port", listenAddr)

	ln, err := net.Listen("tcp", listenAddr)
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
