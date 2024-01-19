package middleware

import (
	"app/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func JWTAuthentication(ctx *fiber.Ctx) error {
	token, ok := ctx.GetReqHeaders()["X-Api-Token"]

	if !ok {
		return fmt.Errorf("unauthorized")
	}

	if err := utils.ParseToken(token[0]); err != nil {
		return fmt.Errorf("unauthorized")
	}

	ctx.Next()

	return nil
}
