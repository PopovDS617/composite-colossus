package middleware

import (
	"app/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func AdminAuth(ctx *fiber.Ctx) error {

	user, err := utils.GetUserFromContext(ctx)

	if err != nil {
		return err
	}

	if !user.IsAdmin {
		return fmt.Errorf("not authorized")
	}

	return ctx.Next()

}
