package api

import (
	"app/types"

	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(ctx *fiber.Ctx) error {

	user := types.User{FirstName: "John",
		LastName: "Johnson"}

	return ctx.JSON(user)
}
func HandleGetUserById(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	return ctx.JSON(map[string]string{"id": id})
}
