package repository

import "withcustomerrorhandling/internal/model"

type CarRepository interface {
	Get(id int) (model.Car, error)
}
