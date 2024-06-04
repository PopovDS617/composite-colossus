package car

import (
	"withcustomerrorhandling/internal/model"
	"withcustomerrorhandling/internal/repository"
)

type CarService struct {
	carRepository repository.CarRepository
}

func NewCarService(repo repository.CarRepository) CarService {
	return CarService{
		carRepository: repo,
	}
}

func (s CarService) Get(id int) (model.Car, error) {

	car, err := s.carRepository.Get(id)

	if err != nil {
		return model.Car{}, err
	}

	return car, nil

}
