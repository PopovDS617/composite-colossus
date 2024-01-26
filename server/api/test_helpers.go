package api

import (
	"app/db"
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testdb struct {
	db.Store
}

func setupDB(t *testing.T) *testdb {

	testDBURI := os.Getenv("MONGO_DB_URI_TEST")
	testDBName := os.Getenv("MONGO_DB_NAME_TEST")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testDBURI))

	if err != nil {
		log.Fatal(err)
	}

	return &testdb{Store: db.Store{
		User:    db.NewMongoUserStore(client, testDBName),
		Hotel:   db.NewMongoHotelStore(client, testDBName),
		Room:    db.NewMongoRoomStore(client, testDBName),
		Booking: db.NewMongoBookingStore(client, testDBName)},
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

func init() {
	if err := godotenv.Load("/app/.env"); err != nil {
		log.Fatal(err)
	}
}
