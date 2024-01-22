package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID     primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	RoomID     primitive.ObjectID `bson:"room_id,omitempty" json:"room_id,omitempty"`
	DateFrom   time.Time          `bson:"date_from,omitempty" json:"date_from,omitempty"`
	DateTill   time.Time          `bson:"date_till,omitempty" json:"date_till,omitempty"`
	NumPersons int                `bson:"num_persons" json:"num_persons"`
}
