package db

import (
	"app/types"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCollection = "users"

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	CreateUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(context.Context, string, *types.User) error
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

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {

	cursor, err := s.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	var users []*types.User

	err = cursor.All(ctx, &users)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *MongoUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {

	res, err := s.collection.InsertOne(ctx, user)

	if err != nil {
		return nil, err
	}

	user.ID = res.InsertedID.(primitive.ObjectID)

	return user, nil

}

func (s *MongoUserStore) DeleteUser(ctx context.Context, id string) error {

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	res, err := s.collection.DeleteOne(ctx, bson.M{"_id": oid})

	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("an error occured during deletion")
	}
	return nil
}
func (s *MongoUserStore) UpdateUser(ctx context.Context, id string, userData *types.User) error {

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}

	updateData := bson.M{"$set": userData}

	res, err := s.collection.UpdateOne(ctx, filter, updateData)

	if err != nil {
		return err
	}

	if res.ModifiedCount == 0 {
		return errors.New("an error occured during update")
	}

	return nil
}
