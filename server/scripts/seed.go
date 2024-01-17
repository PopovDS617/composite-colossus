package main

import (
	"app/db"
	"app/types"
	"app/utils"
	"context"
	"fmt"
	"log"
	"math/rand"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	var hotelsList []types.Hotel
	var roomsList []types.Room

	hotelsCh := make(chan []types.Hotel, 1)
	roomsCh := make(chan []types.Room, 1)

	go utils.ReadAndUnmarshal("/assets/hotels.json", hotelsList, hotelsCh)
	go utils.ReadAndUnmarshal("/assets/rooms.json", roomsList, roomsCh)

	ctx := context.Background()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DB_URI))

	if err != nil {
		log.Fatal(err)
	}

	hotelsList, ok := <-hotelsCh

	if !ok {
		log.Fatal(ok)
	}

	roomsList, ok = <-roomsCh

	if !ok {
		log.Fatal(ok)
	}

	hotelStore := db.NewMongoHotelStore(client, db.DB_NAME)

	if err := hotelStore.Drop(ctx); err != nil {
		log.Fatal(err)
	}

	roomStore := db.NewMongoRoomStore(client, db.DB_NAME)
	if err := roomStore.Drop(ctx); err != nil {
		log.Fatal(err)
	}

	insertedHotelsRaw, err := hotelStore.InsertMultipleHotels(ctx, hotelsList)

	if err != nil {
		log.Fatal(err)
	}

	insertedHotels := []primitive.ObjectID{}

	for _, v := range insertedHotelsRaw {

		insertedHotels = append(insertedHotels, v.(primitive.ObjectID))

	}

	for _, v := range roomsList {

		randomIndex := rand.Intn(len(insertedHotels))

		v.HotelID = insertedHotels[randomIndex]

		_, err = roomStore.InsertRoom(ctx, &v)
		if err != nil {
			log.Fatal(err)
		}

	}

	fmt.Println("db seeding completed!")

}
