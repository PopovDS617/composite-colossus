package api

import (
	"app/db"
	"app/types"

	"github.com/gofiber/fiber/v2"
)

type HotelHandler struct {
	hotelStore db.HotelStore
}

func NewHotelHandler(hotelStore db.HotelStore) *HotelHandler {
	return &HotelHandler{hotelStore: hotelStore}
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

	insertedHotel, err := h.hotelStore.Insert(ctx.Context(), &hotel)

	if err != nil {
		return err
	}

	return ctx.JSON(insertedHotel)
}
