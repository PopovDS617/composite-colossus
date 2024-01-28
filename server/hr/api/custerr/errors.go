package custerr

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(code int, msg string) APIError {
	return APIError{
		Code:    code,
		Message: msg,
	}
}

func InvalidID() APIError {
	return APIError{
		Code:    http.StatusBadRequest,
		Message: "invaid id",
	}
}
func BadRequest() APIError {
	return APIError{
		Code:    http.StatusBadRequest,
		Message: "bad request",
	}
}

func NotFound() APIError {
	return APIError{
		Code:    http.StatusNotFound,
		Message: "not found",
	}
}

func Unauthorized() APIError {
	return APIError{
		Code:    http.StatusUnauthorized,
		Message: "unathorized",
	}
}

func TokenExpired() APIError {
	return APIError{
		Code:    http.StatusUnauthorized,
		Message: "unathorized",
	}
}

func (e APIError) Error() string {
	return e.Message
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	if apiError, ok := err.(APIError); ok {

		return ctx.Status(apiError.Code).JSON(apiError)
	}

	customError := NewError(http.StatusBadRequest, err.Error())

	return ctx.Status(customError.Code).JSON(customError)

}
