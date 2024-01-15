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

const dburi = "mongodb://root:password@localhost:27017/?authSource=admin"
const dbname = "my_db"
const userColl = "users"

var config = fiber.Config{

	ErrorHandler: func(ctx *fiber.Ctx, err error) error {

		return ctx.JSON(map[string]string{"error": err.Error()})

	},
}

func main() {

	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))

	if err != nil {
		log.Fatal(err)
	}

	// handlers
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	app := fiber.New(config)

	apiV1 := app.Group("api/v1")

	apiV1.Get("/users", userHandler.HandleGetUsers)
	apiV1.Get("/users/:id", userHandler.HandleGetUserByID)

	app.Listen(*listenAddr)

}
