package db

import (
	"app/types"
	"app/utils"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const roomCollection = "rooms"

type RoomStore interface {
	Dropper
	Insert(context.Context, *types.Room) (*types.Room, error)
	InsertMultiple(context.Context, []types.Room) error
	GetRooms(context.Context, string) ([]*types.Room, error)
	GetById(context.Context, string) (*types.Room, error)
}

type MongoRoomStore struct {
	client     *mongo.Client
	dbname     string
	collection *mongo.Collection
}

func NewMongoRoomStore(client *mongo.Client, dbname string) *MongoRoomStore {

	return &MongoRoomStore{
		client:     client,
		collection: client.Database(dbname).Collection(roomCollection),
	}
}

func (s *MongoRoomStore) Insert(ctx context.Context, room *types.Room) (*types.Room, error) {

	res, err := s.collection.InsertOne(ctx, room)

	if err != nil {
		return nil, err
	}

	room.ID = res.InsertedID.(primitive.ObjectID)

	// TODO: update hotel "rooms" field

	return room, nil
}

func (s *MongoRoomStore) InsertMultiple(ctx context.Context, rooms []types.Room) error {

	_, err := s.collection.InsertMany(ctx, utils.SliceToInterface[types.Room](rooms))

	if err != nil {
		return err
	}

	return nil
}

func (s *MongoRoomStore) Drop(ctx context.Context) error {
	return s.collection.Drop(ctx)
}

func (s *MongoRoomStore) GetRooms(ctx context.Context, id string) ([]*types.Room, error) {

	var filter interface{}

	if len(id) > 0 {

		oid, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return nil, err
		}
		filter = bson.M{"hotel_id": oid}
	} else {
		filter = bson.M{}
	}
	res, err := s.collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	var rooms []*types.Room

	if err = res.All(ctx, &rooms); err != nil {
		return nil, err
	}

	return rooms, nil
}

func (s *MongoRoomStore) GetById(ctx context.Context, id string) (*types.Room, error) {

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var room types.Room

	err = s.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&room)

	if err != nil {
		return nil, err
	}

	return &room, nil
}
