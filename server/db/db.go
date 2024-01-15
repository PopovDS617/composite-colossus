package db

import "go.mongodb.org/mongo-driver/bson/primitive"

const DB_NAME = "hotel-reservation"

func ToObjectID(s string) primitive.ObjectID {
	objID, err := primitive.ObjectIDFromHex(s)

	if err != nil {
		panic(err)
	}

	return objID
}
