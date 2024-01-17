package db

import (
	"app/types"
	"app/utils"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const hotelCollection = "hotels"

type HotelStore interface {
	Store
	InsertHotel(context.Context, *types.Hotel) (*types.Hotel, error)
	InsertMultipleHotels(context.Context, []types.Hotel) ([]interface{}, error)
}

type MongoHotelStore struct {
	client     *mongo.Client
	dbname     string
	collection *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client, dbname string) *MongoHotelStore {

	return &MongoHotelStore{
		client:     client,
		collection: client.Database(dbname).Collection(hotelCollection),
	}
}

func (s *MongoHotelStore) InsertHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {

	res, err := s.collection.InsertOne(ctx, hotel)

	if err != nil {
		return nil, err
	}

	hotel.ID = res.InsertedID.(primitive.ObjectID)

	return hotel, nil
}

func (s *MongoHotelStore) InsertMultipleHotels(ctx context.Context, hotels []types.Hotel) ([]interface{}, error) {

	res, err := s.collection.InsertMany(ctx, utils.SliceToInterface[types.Hotel](hotels))

	if err != nil {
		return nil, err
	}

	return res.InsertedIDs, nil
}

func (s *MongoHotelStore) Drop(ctx context.Context) error {
	return s.collection.Drop(ctx)
}
