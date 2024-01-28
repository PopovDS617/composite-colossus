package utils

import (
	"app/api/custerr"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SetupFiber() *fiber.App {

	var config = fiber.Config{
		ErrorHandler: custerr.ErrorHandler,
	}

	app := fiber.New(config)
	return app
}

func DecodeResBody[T interface{}](res *http.Response, data T) {
	json.NewDecoder(res.Body).Decode(&data)
}
