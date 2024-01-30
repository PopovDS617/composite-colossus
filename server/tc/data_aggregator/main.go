package main

import (
	"data_aggregator/handlers"
	"data_aggregator/middleware"
	"data_aggregator/service"
	"data_aggregator/store"
	"fmt"
	"net/http"
)

func main() {

	var svc service.Aggregator

	listenAddress := ":9000"

	store := store.NewMemoryStore()

	svc = service.NewInvoiceAggregator(store)
	svc = middleware.NewLogMiddleware(svc)

	makeHTTPTransport(listenAddress, svc)

}

func makeHTTPTransport(port string, svc service.Aggregator) {
	fmt.Println("http transport running on port", port)
	http.HandleFunc("/aggregator", handlers.HandleAggregate(svc))
	http.HandleFunc("/invoice", handlers.HandleGetInvoice(svc))
	http.ListenAndServe(port, nil)
}
