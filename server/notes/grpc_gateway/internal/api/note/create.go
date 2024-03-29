package note

import (
	"context"
	"log"

	"gateway/internal/converter"
	desc "gateway/pkg/note_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	id, err := i.NoteService.Create(ctx, converter.ToNoteInfoFromDesc(req.Info))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted note with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
