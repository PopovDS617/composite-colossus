package api

import (
	"app/db/fixtures"
	"app/types"
	"app/utils"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var roomHandlerTestUser = types.CreateUserParams{
	FirstName: "Mark",
	LastName:  "One",
	Email:     "Mark@mail.com",
	Password:  "Mark_One",
	IsAdmin:   false,
}

var roomHandlerTestHotel = types.Hotel{
	Name:     "Test Hotel",
	Location: "Test Region",
	Rooms:    []primitive.ObjectID{},
	Rating:   5,
}

var roomHandlerTestRoom = types.Room{
	Seaside: true,
	Size:    "small",
	Price:   10,
}

func TestGet(t *testing.T) {
	db := setupDB(t)
	defer db.teardown(t)

	var (
		app = utils.SetupFiber()

		user           = fixtures.AddUser(&db.Store, roomHandlerTestUser)
		hotel          = fixtures.AddHotel(&db.Store, &roomHandlerTestHotel)
		insertedRoom   = fixtures.AddRoom(&db.Store, &roomHandlerTestRoom)
		bookingHandler = NewRoomHandler(&db.Store)
		req            = httptest.NewRequest("GET", "/", nil)
		_              = fixtures.AddBooking(&db.Store, user.ID, insertedRoom.ID, &types.Booking{NumPersons: 2})
		_              = fixtures.AddRoomToHotel(&db.Store, hotel.ID.Hex(), insertedRoom.ID.Hex())
	)

	app.Get("/", bookingHandler.HandleGetAllRooms)

	req.Header.Add("Content-Type", "application/json")

	res, err := app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected to get status code 200 but got %d", res.StatusCode)
	}

	var receivedRoom []*types.Room

	if err := json.NewDecoder(res.Body).Decode(&receivedRoom); err != nil {
		t.Fatal(err)
	}

}
