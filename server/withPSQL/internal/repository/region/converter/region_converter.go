package converter

import (
	"withpsql/internal/model"
	repo "withpsql/internal/repository/region/model"
)

func FromRegionRepoToRegionModel(r *repo.Region) *model.Region {

	return &model.Region{
		ID:   r.ID,
		Name: r.Name,
	}
}
