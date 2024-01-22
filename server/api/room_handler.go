package api

import (
	"app/db"
	"app/types"
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
	DateFrom   time.Time `json:"date_from"`
	DateTill   time.Time `json:"date_till"`
	NumPersons int       `bson:"num_persons" json:"num_persons"`
}

func (p RoomBookingParams) validate() error {
	now := time.Now()

	if now.After(p.DateFrom) || now.After((p.DateTill)) {
		return fmt.Errorf("cannot book a room in the past")
	}

	return nil
}

func NewRoomHandler(store *db.Store) *RoomHandler {
	return &RoomHandler{store: store}
}

func (h *RoomHandler) HandlePostRoomBooking(ctx *fiber.Ctx) error {

	roomID := ctx.Params("id")
	userID := ctx.Locals("user_id").(string)
	var roomBookingParams RoomBookingParams

	if err := ctx.BodyParser(&roomBookingParams); err != nil {
		return err
	}

	if err := roomBookingParams.validate(); err != nil {
		return err
	}

	user, err := h.store.User.GetByID(ctx.Context(), userID)

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
