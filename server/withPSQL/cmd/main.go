package main

import (
	"context"
	"fmt"
	"log"

	"withpsql/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
	fmt.Println("server listening on port 9000")
}
