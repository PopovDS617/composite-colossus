package api

import (
	"app/db"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userStore db.UserStore
}

type AuthParams struct {
	Email    string ` json:"email"`
	Password string `json:"password"`
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
			return fmt.Errorf("invalid credentials")
		}
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(authParams.Password))

	if err != nil {
		return fmt.Errorf("invalid credentials")
	}

	return ctx.JSON(&user)

}