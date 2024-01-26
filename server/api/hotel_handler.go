package api

import (
	"app/api/custerr"
	"app/db"
	"app/types"

	"github.com/gofiber/fiber/v2"
)

type HotelHandler struct {
	store *db.Store
}

type HotelsRes struct {
	Data           []*types.Hotel `json:"results"`
	BatchSize      int            `json:"batch_size"`
	PaginationPage int            `json:"pagination_page"`
}

type HotelsQueryParams struct {
	db.Pagination
	Rating int
}

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{store: store}
}

func (h *HotelHandler) HandleGetHotels(ctx *fiber.Ctx) error {
	var qparams HotelsQueryParams

	if err := ctx.QueryParser(&qparams); err != nil {
		return custerr.BadRequest()
	}

	filter := db.Map{
		"rating": qparams.Rating,
	}

	hotels, err := h.store.Hotel.GetAll(ctx.Context(), filter, &qparams.Pagination)

	if err != nil {
		return custerr.BadRequest()
	}

	response := HotelsRes{
		BatchSize:      int(qparams.BatchSize),
		PaginationPage: int(qparams.Page),
		Data:           hotels,
	}

	return ctx.JSON(response)

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
