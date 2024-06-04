package customerrors

import (
	"errors"
	"net/http"
)

type StatusError struct {
	error
	status int
}

func (e StatusError) Unwrap() error {
	return e.error
}

func (e StatusError) HTTPStatus() int {
	return e.status
}

func WithHTTStatus(error error, status int) StatusError {
	return StatusError{error, status}
}

func HTTPStatus(err error) int {
	if err == nil {
		return 0
	}
	var statusErr interface {
		error
		HTTPStatus() int
	}
	if errors.As(err, &statusErr) {
		status := statusErr.HTTPStatus()
		if status >= 400 && status <= 599 {
			return status
		}
	}

	return http.StatusInternalServerError

}
