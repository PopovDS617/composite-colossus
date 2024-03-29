package main

import (
	"app/api"
	"app/api/custerr"
	"app/api/middleware"
	"app/db"
	"fmt"
	"os"

	"context"

	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	ErrorHandler: custerr.ErrorHandler,
}

func main() {

	MongoDBURI := os.Getenv("MONGO_DB_URI")
	MongoDBName := os.Getenv(db.EnvName)

	listenAddress := os.Getenv("HTTP_LISTEN_ADDRESS")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoDBURI))

	if err != nil {
		log.Fatal(err)
	}

	// stores
	var (
		store = &db.Store{
			Hotel:   db.NewMongoHotelStore(client, MongoDBName),
			Room:    db.NewMongoRoomStore(client, MongoDBName),
			User:    db.NewMongoUserStore(client, MongoDBName),
			Booking: db.NewMongoBookingStore(client, MongoDBName),
		}
	)

	// handlers
	userHandler := api.NewUserHandler(store.User)
	hotelHandler := api.NewHotelHandler(store)
	authHandler := api.NewAuthHandler(store.User)
	roomHandler := api.NewRoomHandler(store)
	bookingHandler := api.NewBookingHandler(store)

	app := fiber.New(config)

	// ROUTER
	apiV1 := app.Group("api/v1")

	// PUBLIC
	// users
	apiV1.Post("/users/", userHandler.HandlePostUser)
	apiV1.Delete("/users/:id", userHandler.HandleDeleteUser)
	apiV1.Patch("/users/:id", userHandler.HandlePatchUser)

	// auth
	apiV1.Post("/login", authHandler.HandleAuth)

	// hotels
	apiV1.Get("/hotels", hotelHandler.HandleGetHotels)
	apiV1.Get("/hotels/:id", hotelHandler.HandleGetHotel)
	/* protected */
	apiV1.Get("/hotels/:id/rooms", middleware.JWTAuthentication(store.User), hotelHandler.HandleGetRoomsByHotelID)

	// rooms
	apiV1.Get("/rooms/", roomHandler.HandleGetAllRooms)
	/* protected */
	apiV1.Post("/rooms/:id/booking", middleware.JWTAuthentication(store.User), roomHandler.HandlePostRoomBooking)

	// booking
	/* protected */
	apiV1.Get("/booking/:id", middleware.JWTAuthentication(store.User), bookingHandler.HandleGetBooking)
	apiV1.Post("/booking/:id", middleware.JWTAuthentication(store.User), bookingHandler.HandleCancelRoomBooking)

	// PRIVATE
	// users
	apiV1.Get("/users", middleware.JWTAuthentication(store.User), middleware.AdminAuth, userHandler.HandleGetUsers)
	apiV1.Get("/users/:id", userHandler.HandleGetUserByID)
	// booking
	apiV1.Get("/booking", middleware.JWTAuthentication(store.User), middleware.AdminAuth, bookingHandler.HandleGetBookings)

	app.Listen(fmt.Sprintf(":%s", listenAddress))

}
