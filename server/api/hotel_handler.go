package api

import (
	"app/api/custerr"
	"app/db"

	"github.com/gofiber/fiber/v2"
)

type HotelHandler struct {
	store *db.Store
}

type GetHotelQueryParams struct {
	Rooms  bool
	Rating int
}

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{store: store}
}

func (h *HotelHandler) HandleGetHotels(ctx *fiber.Ctx) error {

	var qparams GetHotelQueryParams

	if err := ctx.QueryParser(&qparams); err != nil {
		return custerr.BadRequest()
	}

	hotels, err := h.store.Hotel.GetAll(ctx.Context(), "")

	if err != nil {
		return custerr.BadRequest()
	}

	return ctx.JSON(hotels)

}

func (h *HotelHandler) HandleGetRoomsByHotelID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	rooms, err := h.store.Room.GetRooms(ctx.Context(), id)

	if err != nil {
		return custerr.BadRequest()
	}

	return ctx.JSON(rooms)

}

func (h *HotelHandler) HandleGetHotel(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	hotel, err := h.store.Hotel.GetByID(ctx.Context(), id)

	if err != nil {
		return custerr.InvalidID()
	}

	return ctx.JSON(hotel)

}
