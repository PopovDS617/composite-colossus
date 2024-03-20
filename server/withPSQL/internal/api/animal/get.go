package animal

import (
	"context"
	"fmt"
	"net/http"
)

func (i *Implementation) GetAnimalHandler(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// noteObj, err := i.animalService.Get(ctx, req.GetId())
		// if err != nil {
		// 	return nil, err
		// }
		fmt.Println("in get req")
	}
}
