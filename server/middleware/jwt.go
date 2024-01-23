package middleware

import (
	"app/utils"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func JWTAuthentication(ctx *fiber.Ctx) error {
	token, ok := ctx.GetReqHeaders()["X-Api-Token"]

	if !ok {
		return fmt.Errorf("unauthorized")
	}

	claims, err := utils.ValidateToken(token[0])
	if err != nil {
		return fmt.Errorf("unauthorized")
	}

	expires := claims["expires"].(float64)

	tokenTime := int64(expires)

	tokenTimeUNIX := time.Unix(tokenTime, 0)

	if err != nil {
		fmt.Println("could not parse time")
		return fmt.Errorf("token expired")
	}

	if time.Now().Unix() > tokenTimeUNIX.Unix() {
		return fmt.Errorf("token expired")
	}

	userID := claims["user_id"]

	ctx.Locals("user_id", userID)

	return ctx.Next()
}
