package main

import (
	"context"
	"fmt"
	"withredis/internal/handlers"

	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

var baseURL = "https://swapi.dev/api/people/"

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("redis connection established")

	characherHandler := handlers.NewCharacterHandler(client, baseURL)

	http.HandleFunc("/characters/", characherHandler.GetCharacter)

	log.Fatal(http.ListenAndServe(":8090", nil))
}
