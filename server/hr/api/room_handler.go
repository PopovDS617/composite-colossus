package api

import (
	"app/api/custerr"
	"app/db"
	"app/types"
	"app/utils"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
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
		return custerr.BadRequest()
	}

	var roomBookingParams RoomBookingParams

	if err := ctx.BodyParser(&roomBookingParams); err != nil {
		return custerr.BadRequest()
	}

	if err := roomBookingParams.validate(); err != nil {
		return custerr.BadRequest()
	}

	user, err = h.store.User.GetByID(ctx.Context(), user.ID.Hex())

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {

			return custerr.NotFound()
		}

		return custerr.BadRequest()
	}

	room, err := h.store.Room.GetById(ctx.Context(), roomID)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			ctx.SendStatus(http.StatusNotFound)
			return custerr.NotFound()
		}

		return custerr.BadRequest()
	}

	ok, err := h.store.Booking.IsRoomAvailable(ctx.Context(), roomID, roomBookingParams.DateFrom, roomBookingParams.DateTill)

	if err != nil {
		return custerr.BadRequest()
	}

	if !ok {
		err := custerr.NewError(http.StatusBadRequest, fmt.Sprintf("room %s is alreade booked", roomID))

		return ctx.Status(err.Code).JSON(err)
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

func (h *RoomHandler) HandleGetAllRooms(ctx *fiber.Ctx) error {

	rooms, err := h.store.Room.GetRooms(ctx.Context(), "")

	if err != nil {
		return custerr.BadRequest()
	}

	return ctx.JSON(rooms)

}
