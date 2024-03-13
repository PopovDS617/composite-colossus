package animal

import (
	"withpsql/internal/service"
)

type Implementation struct {
	animalService service.AnimalService
}

func NewImplementation(animalService service.AnimalService) *Implementation {
	return &Implementation{
		animalService: animalService,
	}
}
