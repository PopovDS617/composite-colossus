package api

import (
	"app/db"
	"app/types"
	"app/utils"
	"errors"
	"net/http"

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

type GenericResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func authErrorInvalidCred(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusBadRequest).JSON(GenericResponse{Type: "error", Message: "invalid credentials"})
}

func NewAuthHandler(userStore db.UserStore) *AuthHandler {

	return &AuthHandler{
		userStore: userStore,
	}
}

func (h *AuthHandler) HandleAuth(ctx *fiber.Ctx) error {
	var authParams AuthParams

	if err := ctx.BodyParser(&authParams); err != nil {
		return err
	}

	user, err := h.userStore.GetByEmail(ctx.Context(), authParams.Email)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return authErrorInvalidCred(ctx)
		}
		return authErrorInvalidCred(ctx)
	}

	if !utils.IsPasswordValid(authParams.Password, user.EncryptedPassword) {
		return authErrorInvalidCred(ctx)
	}

	jwtToken := utils.CreateToken(user)

	response := AuthResponse{
		User:  user,
		Token: jwtToken,
	}

	return ctx.JSON(response)

}
