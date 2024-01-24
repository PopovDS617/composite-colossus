package db

import (
	"app/types"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const bookingCollection = "bookings"

type BookingStore interface {
	Dropper
	Insert(context.Context, *types.Booking) (*types.Booking, error)
	GetBookings(context.Context, bson.M) ([]*types.Booking, error)
	GetBookingByID(context.Context, string) (*types.Booking, error)
	UpdateBooking(context.Context, string, *types.Booking) error
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

func (s *MongoBookingStore) Drop(ctx context.Context) error {
	return s.collection.Drop(ctx)
}

func (s *MongoBookingStore) Insert(ctx context.Context, booking *types.Booking) (*types.Booking, error) {

	res, err := s.collection.InsertOne(ctx, &booking)

	if err != nil {

		return booking, err
	}

	booking.ID = res.InsertedID.(primitive.ObjectID)

	return booking, nil
}

func (s *MongoBookingStore) GetBookings(ctx context.Context, filter bson.M) ([]*types.Booking, error) {

	cursor, err := s.collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	var bookings []*types.Booking

	if err := cursor.All(ctx, &bookings); err != nil {
		return nil, err
	}

	return bookings, nil

}
func (s *MongoBookingStore) GetBookingByID(ctx context.Context, bookingID string) (*types.Booking, error) {

	oid, err := primitive.ObjectIDFromHex(bookingID)

	if err != nil {
		return nil, err
	}

	var booking types.Booking

	err = s.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&booking)

	if err != nil {
		return nil, err
	}

	return &booking, nil

}

func (s *MongoBookingStore) UpdateBooking(ctx context.Context, bookingID string, updateData *types.Booking) error {

	bookingOID, err := primitive.ObjectIDFromHex(bookingID)

	if err != nil {
		return err
	}

	val := bson.M{
		"$set": updateData,
	}

	_, err = s.collection.UpdateByID(ctx, bookingOID, val)

	if err != nil {
		return err
	}

	return nil
}
