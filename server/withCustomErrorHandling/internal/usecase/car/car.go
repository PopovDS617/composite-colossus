package car

import (
	"encoding/json"
	"net/http"
	"strconv"
	"withcustomerrorhandling/internal/service"
	customerrors "withcustomerrorhandling/pkg/custom_errors"
)

type CarUsecase struct {
	carService service.CarService
}

func NewCarUsecase(service service.CarService) *CarUsecase {
	return &CarUsecase{carService: service}
}

func (u *CarUsecase) Get(w http.ResponseWriter, r *http.Request) error {
	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		return customerrors.WithHTTStatus(err, http.StatusBadRequest)
	}

	car, err := u.carService.Get(id)

	if err != nil {
		return customerrors.WithHTTStatus(err, http.StatusInternalServerError)
	}
	carJSON, err := json.Marshal(car)
	if err != nil {
		return customerrors.WithHTTStatus(err, http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")

	if _, err := w.Write(carJSON); err != nil {
		return customerrors.WithHTTStatus(err, http.StatusInternalServerError)
	}

	return nil
}
