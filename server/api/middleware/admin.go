package middleware

import (
	"app/api/custerr"
	"app/utils"

	"github.com/gofiber/fiber/v2"
)

func AdminAuth(ctx *fiber.Ctx) error {

	user, err := utils.GetUserFromContext(ctx)

	if err != nil {
		return custerr.Unauthorized()
	}

	if !user.IsAdmin {
		return custerr.Unauthorized()
	}

	return ctx.Next()

}
