package types

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hotel struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"` // or json:"_"
	Name     string               `bson:"name" json:"name"`
	Location string               `bson:"location" json:"location"`
	Rooms    []primitive.ObjectID `bson:"rooms" json:"rooms"`
}

type UpdateHotelParams struct {
	Name     string               `bson:"name" json:"name"`
	Location string               `bson:"location" json:"location"`
	Rooms    []primitive.ObjectID `bson:"rooms" json:"rooms"`
}

func (p UpdateHotelParams) ToBSON() bson.M {

	m := bson.M{}

	if len(p.Name) > 0 {
		m["name"] = p.Name
	}

	if len(p.Location) > 0 {
		m["location"] = p.Location
	}
	if len(p.Rooms) > 0 {
		m["rooms"] = p.Rooms
	}

	return m
}
