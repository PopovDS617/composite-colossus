package db

import (
	"app/types"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCollection = "users"

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
}

type MongoUserStore struct {
	client     *mongo.Client
	dbname     string
	collection *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {

	return &MongoUserStore{
		client:     client,
		collection: client.Database(DB_NAME).Collection(userCollection),
	}
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var user types.User

	err = s.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
