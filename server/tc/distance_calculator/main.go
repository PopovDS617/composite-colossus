package main

import (
	"dist_calc/client"
	"dist_calc/consumer"
	"dist_calc/middleware"
	"dist_calc/service"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
	})

	var (
		kafkaTopic         = "obudata"
		service            = service.NewCalculatorService()
		aggregatorEndpoint = os.Getenv("AGGREGATOR_ADDRESS")
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
