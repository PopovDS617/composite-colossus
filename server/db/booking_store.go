package db

import (
	"app/types"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const bookingCollection = "bookings"

type BookingStore interface {
	Dropper
	Insert(context.Context, *types.Booking) (*types.Booking, error)
}

type MongoBookingStore struct {
	client     *mongo.Client
	dbname     string
	collection *mongo.Collection
}

func NewMongoBookingStore(client *mongo.Client, dbname string) *MongoBookingStore {

	return &MongoBookingStore{
		client:     client,
		collection: client.Database(dbname).Collection(bookingCollection),
	}
}

func (s *MongoBookingStore) Insert(ctx context.Context, booking *types.Booking) (*types.Booking, error) {

	res, err := s.collection.InsertOne(ctx, &booking)

	if err != nil {
		fmt.Println("here")
		return booking, err
	}

	booking.ID = res.InsertedID.(primitive.ObjectID)

	return booking, nil
}

func (s *MongoBookingStore) Drop(ctx context.Context) error {
	return s.collection.Drop(ctx)
}
