package api

import (
	"app/api/custerr"
	"app/db"
	"app/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type BookingHandler struct {
	store *db.Store
}

func NewBookingHandler(store *db.Store) *BookingHandler {
	return &BookingHandler{store: store}
}

func (h *BookingHandler) HandleGetBookings(ctx *fiber.Ctx) error {

	bookings, err := h.store.Booking.GetBookings(ctx.Context(), bson.M{})

	if err != nil {
		return custerr.BadRequest()
	}

	return ctx.JSON(bookings)

}
func (h *BookingHandler) HandleGetBooking(ctx *fiber.Ctx) error {

	bookingID := ctx.Params("id")

	booking, err := h.store.Booking.GetBookingByID(ctx.Context(), bookingID)

	if err != nil {
		return custerr.BadRequest()
	}

	user, err := utils.GetUserFromContext(ctx)

	if err != nil {
		return custerr.BadRequest()
	}

	if booking.UserID != user.ID && !user.IsAdmin {
		return custerr.Unauthorized()
	}

	return ctx.JSON(booking)

}

func (h *BookingHandler) HandleCancelRoomBooking(ctx *fiber.Ctx) error {

	bookingID := ctx.Params("id")

	booking, err := h.store.Booking.GetBookingByID(ctx.Context(), bookingID)

	if err != nil {
		return err
	}

	user, err := utils.GetUserFromContext(ctx)

	if err != nil {
		return err
	}

	if booking.UserID != user.ID && !user.IsAdmin {
		return custerr.Unauthorized()
	}

	booking.Cancelled = true

	if err := h.store.Booking.UpdateBooking(ctx.Context(), bookingID, booking); err != nil {
		return err
	}

	return ctx.SendStatus(http.StatusOK)
}
