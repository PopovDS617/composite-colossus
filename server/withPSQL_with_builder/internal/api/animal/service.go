package animal

import (
	"withpsql/internal/service"
)

type AnimalAPI struct {
	animalService service.AnimalService
}

func NewAnimalAPI(animalService service.AnimalService) *AnimalAPI {
	return &AnimalAPI{
		animalService: animalService,
	}
}
