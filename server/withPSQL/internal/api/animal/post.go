package animal

import (
	"context"
	"net/http"
)

func (i *Implementation) CreateAnimalHandler(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		// id, err := i.animalService.Create(ctx, converter.ToNoteInfoFromDesc(req.GetInfo()))
		// if err != nil {
		// 	return nil, err
		// }

		// log.Printf("inserted note with id: %d", id)

		// return &desc.CreateResponse{
		// 	Id: id,
		// }, nil

	}
}
