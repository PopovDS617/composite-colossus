package main

import (
	"context"
	"log"

	"withpsql/internal/app"
	"withpsql/internal/closer"
)

func main() {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}

}
