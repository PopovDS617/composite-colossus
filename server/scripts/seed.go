package main

import (
	"app/db"
	"app/db/fixtures"
	"app/types"
	"app/utils"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	var hotelsList []types.Hotel
	var roomsList []types.Room
	var usersList []types.CreateUserParams

	var wg sync.WaitGroup

	wg.Add(3)

	go utils.ReadAndUnmarshal("/assets/hotels.json", &hotelsList, &wg)
	go utils.ReadAndUnmarshal("/assets/rooms.json", &roomsList, &wg)
	go utils.ReadAndUnmarshal("/assets/users.json", &usersList, &wg)

	wg.Wait()

	adminUser := types.CreateUserParams{
		FirstName: "admin",
		LastName:  "admin",
		Email:     "admin@mail.com",
		Password:  "password",
		IsAdmin:   true,
	}

	ctx := context.Background()

	MongoDBURI := os.Getenv("MONGO_DB_URI")
	MongoDBName := os.Getenv(db.EnvName)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoDBURI))

	if err != nil {
		log.Fatal(err)
	}

	hotelStore := db.NewMongoHotelStore(client, MongoDBName)

	if err := hotelStore.Drop(ctx); err != nil {
		log.Fatal(err)
	}

	roomStore := db.NewMongoRoomStore(client, MongoDBName)
	if err := roomStore.Drop(ctx); err != nil {
		log.Fatal(err)
	}

	userStore := db.NewMongoUserStore(client, MongoDBName)
	if err := userStore.Drop(ctx); err != nil {
		log.Fatal(err)
	}
	bookingStore := db.NewMongoBookingStore(client, MongoDBName)
	if err := bookingStore.Drop(ctx); err != nil {
		log.Fatal(err)
	}

	db := &db.Store{
		User:    userStore,
		Hotel:   hotelStore,
		Room:    roomStore,
		Booking: bookingStore,
	}

	var insertedUsers []*types.User

	insertedAdminUser := fixtures.AddUser(db, adminUser)
	_ = insertedAdminUser

	for _, v := range usersList {

		user := fixtures.AddUser(db, v)

		insertedUsers = append(insertedUsers, user)

	}

	for i, v := range hotelsList {
		hotel := fixtures.AddHotel(db, &v)

		hotelsList[i].ID = hotel.ID
	}

	if err != nil {
		log.Fatal(err)
	}

	for _, currRoom := range roomsList {

		randomHotelIndex := rand.Intn(len(hotelsList))

		randomUserIndex := rand.Intn(len(insertedUsers))

		currRoom.HotelID = hotelsList[randomHotelIndex].ID

		insertedRoom := fixtures.AddRoom(db, &currRoom)

		fixtures.AddRoomToHotel(db, currRoom.HotelID.Hex(), insertedRoom.ID.Hex())

		bookingData := &types.Booking{
			NumPersons: 2,
		}

		fixtures.AddBooking(db, insertedUsers[randomUserIndex].ID, insertedRoom.ID, bookingData)

	}

	fmt.Println("db seeding completed!")

}
