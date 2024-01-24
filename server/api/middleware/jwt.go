package middleware

import (
	"app/api/custerr"
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
			return custerr.Unauthorized()
		}

		claims, err := utils.ValidateToken(token[0])
		if err != nil {
			return custerr.Unauthorized()
		}

		expires := claims["expires"].(float64)

		tokenTime := int64(expires)

		tokenTimeUNIX := time.Unix(tokenTime, 0)

		if err != nil {
			fmt.Println("could not parse time")
			return custerr.TokenExpired()
		}

		if time.Now().Unix() > tokenTimeUNIX.Unix() {
			return custerr.TokenExpired()
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
