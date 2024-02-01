package main

import (
	"dist_calc/client"
	"dist_calc/consumer"
	"dist_calc/middleware"
	"dist_calc/service"
	"log"
)

func main() {
	var (
		kafkaTopic = "obudata"
		service    = service.NewCalculatorService()
		// aggregatorEndpoint = "http://data_aggregator:9000/aggregator"
		aggregatorEndpoint = "http://localhost:9000/aggregator"
	)
	service = middleware.NewLogMiddleware(service)

	httpClient := client.NewHTTPClient(aggregatorEndpoint)

	grpcClient, err := client.NewGRPCClient(":9001")

	if err != nil {
		log.Fatal(err)
	}

	consumer, err := consumer.NewDataConsumer(kafkaTopic, service, httpClient, grpcClient)
	if err != nil {
		log.Fatal(err)
	}

	consumer.Start()
}
