package db

import (
	"app/types"
	"app/utils"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const hotelCollection = "hotels"

type HotelStore interface {
	Insert(context.Context, *types.Hotel) (*types.Hotel, error)
	InsertMultiple(context.Context, []types.Hotel) ([]interface{}, error)
	Update(context.Context, string, *types.UpdateHotelParams) error
	GetByID(context.Context, string) (*types.Hotel, error)
	PushRoom(context.Context, string, string) error
	GetAll(context.Context, string) ([]*types.Hotel, error)
	Dropper
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

func (s *MongoHotelStore) Insert(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {

	res, err := s.collection.InsertOne(ctx, hotel)

	if err != nil {
		return nil, err
	}

	hotel.ID = res.InsertedID.(primitive.ObjectID)

	return hotel, nil
}

func (s *MongoHotelStore) InsertMultiple(ctx context.Context, hotels []types.Hotel) ([]interface{}, error) {

	res, err := s.collection.InsertMany(ctx, utils.SliceToInterface[types.Hotel](hotels))

	if err != nil {
		return nil, err
	}

	return res.InsertedIDs, nil
}

func (s *MongoHotelStore) Update(ctx context.Context, id string, hotelData *types.UpdateHotelParams) error {

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}

	updateData := bson.M{"$set": hotelData.ToBSON()}

	_ = s.collection.FindOneAndUpdate(ctx, filter, updateData)

	if err != nil {
		return err
	}

	return nil
}

func (s *MongoHotelStore) Drop(ctx context.Context) error {
	return s.collection.Drop(ctx)
}

func (s *MongoHotelStore) GetByID(ctx context.Context, id string) (*types.Hotel, error) {

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var hotel types.Hotel

	err = s.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&hotel)

	if err != nil {
		return nil, err
	}

	return &hotel, nil
}

func (s *MongoHotelStore) PushRoom(ctx context.Context, hotelID string, roomID string) error {

	hotelOID, err := primitive.ObjectIDFromHex(hotelID)

	if err != nil {
		return err
	}

	roomOID, err := primitive.ObjectIDFromHex(roomID)

	if err != nil {
		return err
	}

	_, err = s.collection.UpdateOne(ctx, bson.M{"_id": hotelOID}, bson.M{"$push": bson.M{"rooms": roomOID}})

	if err != nil {
		return err
	}

	return nil
}

func (s *MongoHotelStore) GetAll(ctx context.Context, filter string) ([]*types.Hotel, error) {

	res, err := s.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	var hotels []*types.Hotel

	if err = res.All(ctx, &hotels); err != nil {
		return nil, err
	}

	return hotels, err

}
