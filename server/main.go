package main

import (
	"app/api"
	"app/types"
	"context"
	"flag"
	"fmt"

	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://root:password@localhost:27017/my_db?authSource=admin"
const dbname = "my_db"
const userColl = "users"

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	user := types.User{
		FirstName: "Mark",
		LastName:  "One",
	}

	coll := client.Database(dbname).Collection(userColl)
	res, err := coll.InsertOne(ctx, user)

	if err != nil {
		log.Fatal()
	}

	fmt.Println(res)

	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()

	apiV1 := app.Group("api/v1")

	apiV1.Get("/user", api.HandleGetUsers)
	apiV1.Get("/user/:id", api.HandleGetUserById)

	app.Listen(*listenAddr)

}
