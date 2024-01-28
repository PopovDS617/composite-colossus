package api

import (
	"app/api/middleware"
	"app/db/fixtures"
	"app/types"
	"app/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var bookingHandlerTestAdmin = types.CreateUserParams{
	FirstName: "admin",
	LastName:  "admin",
	Email:     "admin@mail.com",
	Password:  "admin_admin",
	IsAdmin:   true,
}

var bookingHandlerTestUser = types.CreateUserParams{
	FirstName: "Mark",
	LastName:  "One",
	Email:     "Mark@mail.com",
	Password:  "Mark_One",
	IsAdmin:   false,
}

var bookingHandlerTestNotAuthUser = types.CreateUserParams{
	FirstName: "Mark",
	LastName:  "Two",
	Email:     "Mark@mail.com",
	Password:  "Mark_Two",
	IsAdmin:   false,
}

var bookingHandlerTestHotel = types.Hotel{
	Name:     "Test Hotel",
	Location: "Test Region",
	Rooms:    []primitive.ObjectID{},
	Rating:   5,
}

var bookingHandlerTestRoom = types.Room{
	Seaside: true,
	Size:    "small",
	Price:   10,
}

func TestAdminGetBookings(t *testing.T) {
	db := setupDB(t)
	defer db.teardown(t)

	var (
		app = utils.SetupFiber()

		admin          = fixtures.AddUser(&db.Store, bookingHandlerTestAdmin)
		user           = fixtures.AddUser(&db.Store, bookingHandlerTestUser)
		hotel          = fixtures.AddHotel(&db.Store, &bookingHandlerTestHotel)
		room           = fixtures.AddRoom(&db.Store, &bookingHandlerTestRoom)
		bookingHandler = NewBookingHandler(&db.Store)
		req            = httptest.NewRequest("GET", "/", nil)
		booking        = fixtures.AddBooking(&db.Store, user.ID, room.ID, &types.Booking{NumPersons: 2})
		_              = fixtures.AddRoomToHotel(&db.Store, hotel.ID.Hex(), room.ID.Hex())
	)

	app.Get("/", middleware.JWTAuthentication(db.Store.User), middleware.AdminAuth, bookingHandler.HandleGetBookings)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Api-Token", utils.CreateToken(admin))

	res, err := app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected to get status code 200 but got %d", res.StatusCode)
	}

	var bookings []*types.Booking

	if err := json.NewDecoder(res.Body).Decode(&bookings); err != nil {
		t.Fatal(err)
	}

	if len(bookings) != 1 {
		t.Fatalf("expected 1 booking got %d", len(bookings))
	}

	have := bookings[0]

	if have.ID != booking.ID {
		t.Fatalf("expected bookings id to be equal")
	}
	if have.RoomID != booking.RoomID {
		t.Fatalf("expected bookings room id to be equal")
	}
	if have.UserID != booking.UserID {
		t.Fatalf("expected bookings user id to be equal")
	}
	if have.NumPersons != booking.NumPersons {
		t.Fatalf("expected bookings number of persons to be equal")
	}
	if have.Cancelled != booking.Cancelled {
		t.Fatalf("expected bookings cancelled status to be equal")
	}

	// non-admin user cannot access the bookings

	req = httptest.NewRequest("GET", "/", nil)
	req.Header.Add("X-Api-Token", utils.CreateToken(user))

	res, err = app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		t.Fatalf("expected to get non 200 status code but got %d", res.StatusCode)
	}

}

func TestUserGetBooking(t *testing.T) {
	db := setupDB(t)
	defer db.teardown(t)

	var (
		app = utils.SetupFiber()

		notAuthUser     = fixtures.AddUser(&db.Store, bookingHandlerTestNotAuthUser)
		user            = fixtures.AddUser(&db.Store, bookingHandlerTestUser)
		hotel           = fixtures.AddHotel(&db.Store, &bookingHandlerTestHotel)
		room            = fixtures.AddRoom(&db.Store, &bookingHandlerTestRoom)
		bookingHandler  = NewBookingHandler(&db.Store)
		req             = httptest.NewRequest("GET", "/:id", nil)
		insertedBooking = fixtures.AddBooking(&db.Store, user.ID, room.ID, &types.Booking{NumPersons: 2})
		_               = fixtures.AddRoomToHotel(&db.Store, hotel.ID.Hex(), room.ID.Hex())
		url             = fmt.Sprintf("/%s", insertedBooking.ID.Hex())
	)

	app.Get("/:id", middleware.JWTAuthentication(db.User), bookingHandler.HandleGetBooking)
	req.Header.Add("Content-Type", "application/json")

	// make req without token
	res, err := app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		t.Fatalf("expected to get non 200 status code but got %d", res.StatusCode)
	}

	// make req with token

	req = httptest.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Api-Token", utils.CreateToken(user))

	res, err = app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected to get status code 200 but got %d", res.StatusCode)
	}

	var receivedBooking *types.Booking

	if err := json.NewDecoder(res.Body).Decode(&receivedBooking); err != nil {
		t.Fatal(err)
	}

	expect := insertedBooking
	have := receivedBooking

	if !reflect.DeepEqual(expect, have) {
		t.Fatal("expected bookings to be equal")
	}

	// make req with token but user is not authorized to get this booking

	req = httptest.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Api-Token", utils.CreateToken(notAuthUser))

	res, err = app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected to get status code 401 but got %d", res.StatusCode)
	}

}
