package car

import (
	"errors"
	"withcustomerrorhandling/internal/model"
)

type repo struct {
}

func NewCarRepository() repo {
	return repo{}
}

func (r repo) Get(id int) (model.Car, error) {
	return model.Car{Price: 250}, errors.New("h")
}
