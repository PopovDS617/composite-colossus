package animal

import (
	"context"
	"net/http"
	"strconv"
	"withpsql/internal/utils"
)

func (api *AnimalAPI) DeleteAnimalHandler(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		animalID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)

		if err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "wrong id format"})
			return
		}

		err = api.animalService.Delete(ctx, int64(animalID))

		if err != nil {
			utils.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
			return
		}

		err = utils.WriteJSON(w, http.StatusOK, nil)

		if err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal error"})
			return
		}

	}
}
