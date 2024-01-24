package middleware

import (
	"app/db"
	"app/utils"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func JWTAuthentication(s db.UserStore) fiber.Handler {

	return func(ctx *fiber.Ctx) error {
		token := ctx.GetReqHeaders()["X-Api-Token"]

		if len(token) == 0 {
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

		userID := claims["user_id"].(string)

		user, err := s.GetByID(ctx.Context(), userID)

		if err != nil {
			return err
		}

		ctx.Context().SetUserValue("user", user)

		return ctx.Next()
	}
}
