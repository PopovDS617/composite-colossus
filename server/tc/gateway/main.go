package main

import (
	"fmt"
	"gateway/client"
	"gateway/handlers"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
	})

	var (
		port              = os.Getenv("PORT")
		aggregatorAddress = os.Getenv("AGGREGATOR_ADDRESS")
		httpClient        = client.NewHTTPClient(aggregatorAddress)
		invHandler        = handlers.NewInvoiceHandler(httpClient)
	)

	http.HandleFunc("/invoice", handlers.MakeAPIFunc(invHandler.HandleGetInvoice))
	logrus.Info("gateway HTTP server running and listening on port ", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
