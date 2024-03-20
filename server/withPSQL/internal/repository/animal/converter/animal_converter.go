package converter

import (
	"withpsql/internal/model"
	repo "withpsql/internal/repository/animal/model"
)

func FromAnimalRepoToAnimalModel(r *repo.Animal) *model.Animal {

	return &model.Animal{
		ID:             r.ID,
		Name:           r.Name,
		Age:            r.Age,
		Gender:         r.Gender,
		Type:           r.Type,
		CreatedAt:      r.CreatedAt,
		UpdatedAt:      r.UpdatedAt,
		Region:         r.RegionName,
		LastTimeSeenAt: r.LastTimeSeenAt,
		SeenByDevice:   r.SeenByDevice,
	}
}
