package handlers

import (
	"net/http"
	customerrors "withcustomerrorhandling/pkg/custom_errors"
)

// customHandler преобразует обработчик, возвращающий ошибку, в стандартный http.HandlerFunc.
func CustomHandler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)

		if err != nil {
			http.Error(w, err.Error(), customerrors.HTTPStatus(err))
		}
	}
}
