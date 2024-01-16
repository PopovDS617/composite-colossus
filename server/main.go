package main

import (
	"app/api"
	"app/db"
	"context"
	"flag"

	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{

	ErrorHandler: func(ctx *fiber.Ctx, err error) error {

		return ctx.JSON(map[string]string{"error": err.Error()})

	},
}

func main() {

	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DB_URI))

	if err != nil {
		log.Fatal(err)
	}

	// handlers
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client, db.DB_NAME))

	app := fiber.New(config)

	apiV1 := app.Group("api/v1")

	apiV1.Get("/users", userHandler.HandleGetUsers)
	apiV1.Get("/users/:id", userHandler.HandleGetUserByID)
	apiV1.Post("/users/", userHandler.HandlePostUser)
	apiV1.Delete("/users/:id", userHandler.HandleDeleteUser)
	apiV1.Patch("/users/:id", userHandler.HandlePatchUser)

	app.Listen(*listenAddr)

}
