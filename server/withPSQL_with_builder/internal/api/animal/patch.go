package animal

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"withpsql/internal/model"
	"withpsql/internal/utils"
)

func (api *AnimalAPI) UpdateAnimalHandler(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		animalID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)

		if err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "wrong id format"})
			return
		}
		var receivedAnimal model.Animal

		if err := json.NewDecoder(r.Body).Decode(&receivedAnimal); err != nil {
			if errors.Is(err, io.EOF) {
				utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "request body is empty"})
				return
			}
			utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "wrong request body format"})
			return
		}
		defer r.Body.Close()

		storedAnimal, err := api.animalService.Get(ctx, animalID)

		if err != nil {
			utils.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
			return
		}

		storedAnimal.ValidateAndUpdate(&receivedAnimal)

		updatedAnimal, err := api.animalService.Update(ctx, storedAnimal)
		if err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, nil)
			return
		}

		utils.WriteJSON(w, http.StatusOK, updatedAnimal)
	}
}
