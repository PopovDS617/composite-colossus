package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Room struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Seaside bool               `bson:"seaside" json:"seaside"`
	Size    string             `bson:"size" json:"size"`
	Price   float64            `bson:"price" json:"price"`
	HotelID primitive.ObjectID `bson:"hotel_id,omitempty" json:"hotel_id,omitempty"`
}
