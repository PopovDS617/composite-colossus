package db

import (
	"app/types"
	"app/utils"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const roomCollection = "rooms"

type RoomStore interface {
	Store
	InsertRoom(context.Context, *types.Room) (*types.Room, error)
	InsertMultipleRooms(context.Context, []types.Room) error
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

func (s *MongoRoomStore) InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error) {

	res, err := s.collection.InsertOne(ctx, room)

	if err != nil {
		return nil, err
	}

	room.ID = res.InsertedID.(primitive.ObjectID)

	return room, nil
}

func (s *MongoRoomStore) InsertMultipleRooms(ctx context.Context, rooms []types.Room) error {

	_, err := s.collection.InsertMany(ctx, utils.SliceToInterface[types.Room](rooms))

	if err != nil {
		return err
	}

	return nil
}

func (s *MongoRoomStore) Drop(ctx context.Context) error {
	return s.collection.Drop(ctx)
}
