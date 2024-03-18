package animal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"withpsql/internal/model"
	"withpsql/internal/utils"
)

func (i *Implementation) CreateAnimalHandler(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		var animal model.Animal

		if err := json.NewDecoder(r.Body).Decode(&animal); err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, nil)
			return
		}

		inputErrorsMap, ok := animal.ValidateCreate()

		if !ok {
			utils.WriteJSON(w, http.StatusBadRequest, inputErrorsMap)
			return
		}

		insertedAnimal, err := i.animalService.Create(ctx, &animal)
		if err != nil {
			fmt.Println(err)
			utils.WriteJSON(w, http.StatusBadRequest, nil)
			return
		}

		utils.WriteJSON(w, http.StatusCreated, insertedAnimal)
	}
}
