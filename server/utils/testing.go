package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SetupFiber() *fiber.App {
	app := fiber.New()
	return app
}

func DecodeResBody[T interface{}](res *http.Response, data T) {
	json.NewDecoder(res.Body).Decode(&data)
}
