package api

import (
	"app/api/custerr"
	"app/db"

	"app/types"
	"app/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthHandler struct {
	userStore db.UserStore
}

type AuthParams struct {
	Email    string ` json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	User  *types.User `json:"user"`
	Token string      `json:"token"`
}

func NewAuthHandler(userStore db.UserStore) *AuthHandler {

	return &AuthHandler{
		userStore: userStore,
	}
}

func (h *AuthHandler) HandleAuth(ctx *fiber.Ctx) error {
	var authParams AuthParams

	if err := ctx.BodyParser(&authParams); err != nil {
		return custerr.BadRequest()
	}

	user, err := h.userStore.GetByEmail(ctx.Context(), authParams.Email)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return custerr.BadRequest()
		}
		return custerr.BadRequest()
	}

	if !utils.IsPasswordValid(authParams.Password, user.EncryptedPassword) {
		return custerr.BadRequest()
	}

	jwtToken := utils.CreateToken(user)

	response := AuthResponse{
		User:  user,
		Token: jwtToken,
	}

	return ctx.JSON(response)

}
