package region

import (
	"withpsql/internal/service"
)

type RegionAPI struct {
	regionService service.RegionService
}

func NewRegionAPI(regionService service.RegionService) *RegionAPI {
	return &RegionAPI{
		regionService: regionService,
	}
}
