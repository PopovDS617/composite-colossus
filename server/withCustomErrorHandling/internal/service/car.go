package service

import "withcustomerrorhandling/internal/model"

type CarService interface {
	Get(id int) (model.Car, error)
}
