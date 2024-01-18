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
	Dropper
	GetByID(context.Context, string) (*types.User, error)
	GetAll(context.Context) ([]*types.User, error)
	Insert(context.Context, *types.User) (*types.User, error)
	Delete(context.Context, string) error
	Update(context.Context, string, *types.UpdateUserParams) error
}

type MongoUserStore struct {
	client     *mongo.Client
	dbname     string
	collection *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client, dbname string) *MongoUserStore {

	return &MongoUserStore{
		client:     client,
		collection: client.Database(dbname).Collection(userCollection),
	}
}

func (s *MongoUserStore) GetByID(ctx context.Context, id string) (*types.User, error) {

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

func (s *MongoUserStore) GetAll(ctx context.Context) ([]*types.User, error) {

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

func (s *MongoUserStore) Insert(ctx context.Context, user *types.User) (*types.User, error) {

	res, err := s.collection.InsertOne(ctx, user)

	if err != nil {
		return nil, err
	}

	user.ID = res.InsertedID.(primitive.ObjectID)

	return user, nil

}

func (s *MongoUserStore) Delete(ctx context.Context, id string) error {

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
func (s *MongoUserStore) Update(ctx context.Context, id string, userData *types.UpdateUserParams) error {

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}

	updateData := bson.M{"$set": userData.ToBSON()}

	res, err := s.collection.UpdateOne(ctx, filter, updateData)

	if err != nil {
		return err
	}

	if res.ModifiedCount == 0 {
		return errors.New("an error occured during update")
	}

	return nil
}

func (s *MongoUserStore) Drop(ctx context.Context) error {
	return s.collection.Drop(ctx)
}
