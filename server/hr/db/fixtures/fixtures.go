package fixtures

import (
	"app/db"
	"app/types"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUser(store *db.Store, preparedUserData types.CreateUserParams) *types.User {

	user, err := types.NewUserFromParams(preparedUserData)

	if err != nil {
		log.Fatal(err)
	}

	insertedUser, err := store.User.Insert(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}

	return insertedUser
}

func AddHotel(store *db.Store, preparedHotelData *types.Hotel) *types.Hotel {

	insertedHotel, err := store.Hotel.Insert(context.Background(), preparedHotelData)

	if err != nil {
		log.Fatal(err)
	}

	return insertedHotel
}

func AddRoom(store *db.Store, preparedRoomData *types.Room) *types.Room {

	insertedRoom, err := store.Room.Insert(context.Background(), preparedRoomData)

	if err != nil {
		log.Fatal(err)
	}

	return insertedRoom
}

func AddRoomToHotel(store *db.Store, hotelID, roomID string) error {

	store.Hotel.PushRoom(context.Background(), hotelID, roomID)

	return nil

}

func AddBooking(store *db.Store, userID, roomID primitive.ObjectID, bookingData *types.Booking) *types.Booking {
	parsedTimeFrom, _ := time.Parse(time.RFC3339, bookingData.DateFrom)

	parsedTimeTill, _ := time.Parse(time.RFC3339, bookingData.DateTill)

	if parsedTimeFrom.IsZero() {
		bookingData.DateFrom = time.Now().AddDate(0, 0, 5).UTC().Format(time.RFC3339)
	}
	if parsedTimeTill.IsZero() {
		bookingData.DateTill = time.Now().AddDate(0, 0, 12).UTC().Format(time.RFC3339)
	}

	if bookingData.NumPersons == 0 {
		bookingData.NumPersons = 2
	}

	booking := &types.Booking{
		UserID:     userID,
		RoomID:     roomID,
		NumPersons: bookingData.NumPersons,
		DateFrom:   bookingData.DateFrom,
		DateTill:   bookingData.DateTill,
	}

	insertedBooking, _ := store.Booking.Insert(context.Background(), booking)

	return insertedBooking
}
