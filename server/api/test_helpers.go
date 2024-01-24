package api

import (
	"app/db"
	"context"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testdb struct {
	db.Store
}

const testdburi = "mongodb://root:password@localhost:27017/?authSource=admin"
const testdbname = "hotel-reservation-test"

func setupDB(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testdburi))

	if err != nil {
		log.Fatal(err)
	}

	return &testdb{Store: db.Store{
		User:    db.NewMongoUserStore(client, testdbname),
		Hotel:   db.NewMongoHotelStore(client, testdbname),
		Room:    db.NewMongoRoomStore(client, testdbname),
		Booking: db.NewMongoBookingStore(client, testdbname)},
	}

}

func (tdb *testdb) teardown(t *testing.T) {

	var err error

	err = tdb.User.Drop(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	err = tdb.Hotel.Drop(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	err = tdb.Room.Drop(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	err = tdb.Booking.Drop(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
}
