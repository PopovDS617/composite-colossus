package main

import (
	"log"
	"withcustomerrorhandling/internal/app"
)

func main() {
	app := app.NewApp()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
