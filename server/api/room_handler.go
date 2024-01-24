package api

import (
	"app/db"
	"app/types"
	"app/utils"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomHandler struct {
	store *db.Store
}

type RoomBookingParams struct {
	DateFrom   string `json:"date_from"`
	DateTill   string `json:"date_till"`
	NumPersons int    `bson:"num_persons" json:"num_persons"`
}

func (p RoomBookingParams) validate() error {
	now := time.Now()

	parsedTimeFrom, err := time.Parse(time.RFC3339, p.DateFrom)

	if err != nil {
		return fmt.Errorf("wrong date format")
	}
	parsedTimeTill, err := time.Parse(time.RFC3339, p.DateTill)

	if err != nil {
		return fmt.Errorf("wrong date format")
	}

	if now.After(parsedTimeFrom) || now.After(parsedTimeTill) {
		return fmt.Errorf("cannot book a room in the past")
	}

	return nil
}

func NewRoomHandler(store *db.Store) *RoomHandler {
	return &RoomHandler{store: store}
}

func (h *RoomHandler) HandlePostRoomBooking(ctx *fiber.Ctx) error {

	roomID := ctx.Params("id")
	user, err := utils.GetUserFromContext(ctx)

	if err != nil {
		return err
	}

	var roomBookingParams RoomBookingParams

	if err := ctx.BodyParser(&roomBookingParams); err != nil {
		return err
	}

	if err := roomBookingParams.validate(); err != nil {
		return err
	}

	user, err = h.store.User.GetByID(ctx.Context(), user.ID.Hex())

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			ctx.SendStatus(http.StatusNotFound)
			return ctx.JSON(map[string]string{"message": "not found"})
		}

		return err
	}

	room, err := h.store.Room.GetById(ctx.Context(), roomID)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			ctx.SendStatus(http.StatusNotFound)
			return ctx.JSON(map[string]string{"message": "not found"})
		}

		return err
	}

	ok, err := h.isRoomAvailable(ctx.Context(), roomID, &roomBookingParams)

	if err != nil {
		return err
	}

	if !ok {
		return ctx.Status(http.StatusBadRequest).JSON(GenericResponse{
			Type:    "error",
			Message: fmt.Sprintf("room %s is alreade booked", roomID),
		})
	}

	booking := types.Booking{
		UserID:     user.ID,
		RoomID:     room.ID,
		DateFrom:   roomBookingParams.DateFrom,
		DateTill:   roomBookingParams.DateTill,
		NumPersons: roomBookingParams.NumPersons,
	}

	insertedBooking, err := h.store.Booking.Insert(ctx.Context(), &booking)

	if err != nil {
		return err
	}

	return ctx.JSON(insertedBooking)
}

func (h *RoomHandler) isRoomAvailable(ctx context.Context, roomID string, roomBookingParams *RoomBookingParams) (bool, error) {

	roomOID, err := primitive.ObjectIDFromHex(roomID)

	if err != nil {
		return false, err
	}

	filter := bson.M{
		"room_id": roomOID,
		"date_from": bson.M{
			"$gte": roomBookingParams.DateFrom,
		},
		"date_till": bson.M{
			"$lte": roomBookingParams.DateTill,
		},
	}

	bookings, err := h.store.Booking.GetBookings(ctx, filter)

	if err != nil {
		return false, err
	}

	ok := len(bookings) == 0
	return ok, nil
}

func (h *RoomHandler) HandleGetAllRooms(ctx *fiber.Ctx) error {

	rooms, err := h.store.Room.GetRooms(ctx.Context(), "")

	if err != nil {
		return err
	}

	return ctx.JSON(rooms)

}
