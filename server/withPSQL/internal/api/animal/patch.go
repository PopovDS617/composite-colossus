package animal

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"withpsql/internal/model"
	"withpsql/internal/utils"
)

func (i *Implementation) UpdateAnimalHandler(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		animalID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)

		if err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "wrong id format"})
			return
		}
		var animal model.Animal

		if err := json.NewDecoder(r.Body).Decode(&animal); err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, nil)
			return
		}

		animal.ID = animalID

		updatedAnimal, err := i.animalService.Update(ctx, &animal)
		if err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, nil)
			return
		}

		utils.WriteJSON(w, http.StatusCreated, updatedAnimal)
	}
}
