package region

import (
	"context"
	"net/http"
	"withpsql/internal/utils"
)

// func (api *RegionAPI) GetRegionHandler(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		animalID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)

// 		if err != nil {
// 			utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "wrong id format"})
// 			return
// 		}

// 		animal, err := api.animalService.Get(ctx, int64(animalID))

// 		if err != nil {
// 			utils.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
// 			return
// 		}

// 		err = utils.WriteJSON(w, http.StatusOK, animal)

// 		if err != nil {
// 			utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal error"})
// 			return
// 		}

// 	}
// }

func (api *RegionAPI) GetAllRegionsHandler(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		regions, err := api.regionService.GetAll(ctx)

		if err != nil {
			utils.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
			return
		}

		err = utils.WriteJSON(w, http.StatusOK, regions)

		if err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal error"})
			return
		}

	}
}
