package api

import (
	"app/db"
	"app/types"

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

func (h *HotelHandler) HandlePostHotel(ctx *fiber.Ctx) error {
	// var params types.CreateUserParams

	// if err := ctx.BodyParser(&params); err != nil {
	// 	return err
	// }

	// if errList := params.ValidateUser(); len(errList) > 0 {
	// 	return ctx.JSON(errList)
	// }

	// user, err := types.NewUserFromParams(params)

	// if err != nil {
	// 	return err
	// }

	hotel := types.Hotel{
		Name:     "Bellucia",
		Location: "France",
	}

	insertedHotel, err := h.store.Hotel.Insert(ctx.Context(), &hotel)

	if err != nil {
		return err
	}

	return ctx.JSON(insertedHotel)
}

func (h *HotelHandler) HandleGetHotels(ctx *fiber.Ctx) error {

	var qparams GetHotelQueryParams

	if err := ctx.QueryParser(&qparams); err != nil {
		return err
	}

	hotels, err := h.store.Hotel.GetAll(ctx.Context(), "")

	if err != nil {
		return err
	}

	return ctx.JSON(hotels)

}

func (h *HotelHandler) HandleGetRooms(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	rooms, err := h.store.Room.GetRooms(ctx.Context(), id)

	if err != nil {
		return err
	}

	return ctx.JSON(rooms)

}

func (h *HotelHandler) HandleGetHotel(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	hotel, err := h.store.Hotel.GetByID(ctx.Context(), id)

	if err != nil {
		return err
	}

	return ctx.JSON(hotel)

}
