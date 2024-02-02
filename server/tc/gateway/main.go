package main

import (
	"fmt"
	"gateway/client"
	"gateway/handlers"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
	})

	var (
		listenAdderess = "10000"
		httpClient     = client.NewHTTPClient("http://localhost:9000")
		invHandler     = handlers.NewInvoiceHandler(httpClient)
	)

	http.HandleFunc("/invoice", handlers.MakeAPIFunc(invHandler.HandleGetInvoice))
	logrus.Info("gateway HTTP server running and listening on port ", listenAdderess)
	http.ListenAndServe(fmt.Sprintf(":%v", listenAdderess), nil)
}
