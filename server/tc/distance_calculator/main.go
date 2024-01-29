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
		kafkaTopic         = "obudata"
		service            = service.NewCalculatorService()
		aggregatorEndpoint = "http://localhost:3000/aggregator"
	)
	service = middleware.NewLogMiddleware(service)

	client := client.NewClient(aggregatorEndpoint)

	consumer, err := consumer.NewDataConsumer(kafkaTopic, service, client)

	if err != nil {
		log.Fatal(err)
	}

	consumer.Start()
}
