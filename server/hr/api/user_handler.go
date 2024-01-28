package api

import (
	"app/api/custerr"
	"app/db"
	"app/types"
	"errors"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{userStore: userStore}
}

func (h *UserHandler) HandleGetUsers(ctx *fiber.Ctx) error {

	users, err := h.userStore.GetAll(ctx.Context())

	if err != nil {
		return custerr.BadRequest()
	}

	return ctx.JSON(users)
}
func (h *UserHandler) HandleGetUserByID(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	user, err := h.userStore.GetByID(ctx.Context(), id)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return custerr.NotFound()
		}

		return err
	}

	return ctx.JSON(user)
}

func (h *UserHandler) HandlePostUser(ctx *fiber.Ctx) error {
	var params types.CreateUserParams

	if err := ctx.BodyParser(&params); err != nil {
		return custerr.BadRequest()
	}

	if errList := params.ValidateUser(); len(errList) > 0 {
		return ctx.JSON(errList)
	}

	user, err := types.NewUserFromParams(params)

	if err != nil {
		return custerr.BadRequest()
	}

	insertedUser, err := h.userStore.Insert(ctx.Context(), user)

	if err != nil {
		return custerr.BadRequest()
	}

	return ctx.JSON(insertedUser)
}

func (h *UserHandler) HandleDeleteUser(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")

	err := h.userStore.Delete(ctx.Context(), userID)

	if err != nil {
		return custerr.BadRequest()
	}

	return ctx.JSON(map[string]string{"message": "deleted successfully"})
}

func (h *UserHandler) HandlePatchUser(ctx *fiber.Ctx) error {

	userID := ctx.Params("id")

	var updatedUserData types.UpdateUserParams

	err := ctx.BodyParser(&updatedUserData)

	if err != nil {
		return custerr.BadRequest()
	}

	err = h.userStore.Update(ctx.Context(), userID, &updatedUserData)

	if err != nil {
		return custerr.BadRequest()
	}

	return ctx.JSON(map[string]string{"message": "successfully updated"})
}
