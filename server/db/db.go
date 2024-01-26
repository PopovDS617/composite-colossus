package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const DB_NAME = "hotel-reservation"

const DB_URI = "mongodb://root:password@localhost:27017/?authSource=admin"

func ToObjectID(s string) primitive.ObjectID {
	objID, err := primitive.ObjectIDFromHex(s)

	if err != nil {
		panic(err)
	}

	return objID
}

type Dropper interface {
	Drop(context.Context) error
}

type Store struct {
	User    UserStore
	Hotel   HotelStore
	Room    RoomStore
	Booking BookingStore
}

type Pagination struct {
	BatchSize int64 `query:"batch_size"`
	Page      int64 `query:"pagination_page"`
}

type Map map[string]interface{}
