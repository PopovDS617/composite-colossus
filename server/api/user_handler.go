package api

import (
	"app/db"
	"app/types"
	"context"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{userStore: userStore}
}

func (h *UserHandler) HandleGetUsers(ctx *fiber.Ctx) error {

	user := types.User{FirstName: "John",
		LastName: "Johnson"}

	return ctx.JSON(user)
}
func (h *UserHandler) HandleGetUserByID(ctx *fiber.Ctx) error {

	var (
		id       = ctx.Params("id")
		ctxInner = context.Background()
	)

	user, err := h.userStore.GetUserByID(ctxInner, id)

	if err != nil {
		return err
	}

	return ctx.JSON(user)
}
